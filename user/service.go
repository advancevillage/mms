//author: richard
package user

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/api"
)

func NewService(storage storages.StorageExd, logger logs.Logs, cache caches.ICache) *Service {
	return &Service{
		logger:logger,
		cache: cache,
		repo: NewRepoMongo(storage),
	}
}

func (s *Service) LoginToken(username string) (string, error) {
	if !utils.ValidateEmail(username) {
		return "", errors.New("invalid email format")
	}
	u, err := s.repo.QueryUserByName(username)
	if err != nil {
		s.logger.Error(err.Error())
		return "", errors.New("your email is non-members, please sign up")
	}
	timestamp := times.Timestamp()
	random    := utils.RandsNumberString(6)
	token     := fmt.Sprintf("%d%s", timestamp, random)
	key, err  := s.CreateKey("login", u.Username, token)
	if err !=nil {
		return "", errors.New("cache key create error")
	}
	err = s.cache.CreateCache(key, []byte(fmt.Sprintf("%d", timestamp)), LoginTimeout)
	if err != nil {
		s.logger.Error(err.Error())
		return "", errors.New("cache key create error")
	}
	return token, nil
}

func (s *Service) RegisterToken(username string) (string, error) {
	if !utils.ValidateEmail(username) {
		return "", errors.New("invalid email format")
	}
	u, err := s.repo.QueryUserByName(username)
	if u != nil {
		return "", errors.New("your email existed, please login")
	}
	timestamp := times.Timestamp()
	random    := utils.RandsNumberString(6)
	token     := fmt.Sprintf("%d%s", timestamp, random)
	key, err  := s.CreateKey("register", username, token)
	if err !=nil {
		return "", errors.New("cache key create error")
	}
	//生成 key
	err = s.cache.CreateCache(key, []byte(fmt.Sprintf("%d", timestamp)), RegisterTimeout)
	if err != nil {
		s.logger.Error(err.Error())
		return "", errors.New("cache key create error")
	}

	return token, nil
}

func (s *Service) QueryUser(login *api.Login) (*api.User, error) {
	//邮箱 有效性
	if !utils.ValidateEmail(login.Username) {
		return nil, errors.New("invalid email format")
	}
	//密码 有效性
	if len(login.Password) != SHA1 >> 2 {
		return nil, errors.New("invalid password")
	}
	//签名 有效性
	if len(login.Sign) != SHA1 >> 2 {
		return nil, errors.New("invalid sign")
	}
	//令牌 有效性
	key, err := s.CreateKey("login", login.Username, login.Token)
	if err != nil {
		return nil, errors.New("cache key create error")
	}
	_, err = s.cache.QueryCache(key, LoginTimeout)
	if err != nil {
		return nil, errors.New("token invalid and re login")
	}
	//无篡改 检查
	h := sha1.New()
	_, err = h.Write([]byte(fmt.Sprintf("%s%s%s%s",login.Username, login.Token, login.Timestamp, login.Password)))
	if err != nil {
		return nil, errors.New("security check failed")
	}
	if login.Sign != fmt.Sprintf("%x", h.Sum(nil)) {
		return nil, errors.New("risk of hijacking and reject to login")
	}
	//查询 用户基本数据
	u, err := s.repo.QueryUserByName(login.Username)
	if err != nil {
		s.logger.Error(err.Error())
		return nil , errors.New("your email does not exist")
	}

	//校验密码
	if login.Password != u.Password {
		return nil , errors.New("wrong password")
	}

	//查询完成 删除key
	err = s.cache.DeleteCache(key)
	if err != nil {
		s.logger.Warning(err.Error())
	}

	return u, nil
}

func (s *Service) CreateUser(register *api.Register) (*api.User, error) {
	//邮箱 有效性
	if !utils.ValidateEmail(register.Username) {
		return nil, errors.New("invalid email format")
	}
	//密码 有效性
	if len(register.Password) != SHA1 >> 2 {
		return nil, errors.New("invalid password")
	}
	//令牌 有效性
	key, err := s.CreateKey("register", register.Username, register.Token)
	if err != nil {
		return nil, errors.New("cache key create error")
	}
	_, err = s.cache.QueryCache(key, RegisterTimeout)
	if err != nil {
		return nil, errors.New("token invalid and re sign")
	}
	//无篡改 检查
	h := sha1.New()
	_, err = h.Write([]byte(fmt.Sprintf("%s%s%s%s",register.Username, register.Token, register.Timestamp, register.Password)))
	if err != nil {
		return nil, errors.New("security check failed")
	}
	if register.Sign != fmt.Sprintf("%x", h.Sum(nil)) {
		return nil, errors.New("risk of hijacking and reject to sign")
	}

	//完成校验 删除key 保持函数幂等性 防止重复创建
	err = s.cache.DeleteCache(key)
	if err != nil {
		s.logger.Warning(err.Error())
	}

	//查询 用户
	u, err := s.repo.QueryUserByName(register.Username)
	if u != nil {
		return nil, errors.New("your email existed")
	}

	//设置 用户基本信息
	user := api.User{}
	user.Id = utils.SnowFlakeIdString()
	user.Gender   = register.Gender
	user.Username = register.Username
	user.Password = register.Password
	user.CreateTime = times.Timestamp()
	user.UpdateTime = times.Timestamp()
	user.DeleteTime = 0

	err = s.repo.CreateUser(&user)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	//TODO 异步 发送第一封推送邮件
	go func() {

	}()
	return &user, nil
}

func (s *Service) CreateKey(str ... string) (string, error) {
	if len(str) <= 0 {
		return "", errors.New("empty key")
	}
	if len(str) == 1 {
		return str[0], nil
	}
	key := str[0]
	for i := 1; i < len(str); i++ {
		key = fmt.Sprintf("%s:%s", key, str[i])
	}
	return key, nil
}

func (s *Service) QueryCart(user *api.User) ([]api.Cart, int, error) {
	if user == nil {
		return nil, 0, errors.New("user is nil")
	}
	carts, total, err := s.repo.QueryCart(user)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	itemTotal := 0
	for i := 0; i < len(carts) && i < int(total); i++ {
		itemTotal += carts[i].Total
	}

	return carts, itemTotal, nil
}

func (s *Service) CreateCart(user *api.User, cart *api.Cart) error {
	if user == nil || cart == nil {
		return errors.New("user or cart is nil")
	}

	carts, total, err := s.QueryCart(user)
	if err != nil {
		s.logger.Error(err.Error())
		return  err
	}

	i := 0
	for ; i < len(carts) && i < total; i++ {
		if carts[i].GoodsId == cart.GoodsId && carts[i].SizeId == cart.SizeId && carts[i].ColorId == carts[i].ColorId {
			carts[i].Total += cart.Total
			break
		} else {
			continue
		}

	}

	if i < len(carts) {
		//购物车已存在该商品则合并购物车
		err = s.repo.UpdateCart(user, &carts[i])
		if err != nil {
			s.logger.Error(err.Error())
			return err
		} else {
			return nil
		}
	}

	cart.Id =  utils.SnowFlakeIdString()
	cart.CreateTime = times.Timestamp()
	cart.UpdateTime = times.Timestamp()
	cart.DeleteTime = 0
	err = s.repo.CreateCart(user, cart)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateCart(user *api.User, cart *api.Cart) error {
	if user == nil || cart == nil {
		return errors.New("user or cart is nil")
	}
	cart.UpdateTime = times.Timestamp()
	err := s.repo.UpdateCart(user, cart)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteCart(user *api.User, cartId string) error {
	if user == nil {
		return errors.New("user is nil")
	}
	cart, err := s.QueryOneCart(user, cartId)
	if err != nil {
		return err
	}
	cart.DeleteTime = times.Timestamp()
	cart.UpdateTime = times.Timestamp()
	err = s.repo.DeleteCart(user, cart)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) QueryOneCart(user *api.User, cartId string) (*api.Cart, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	cart, err := s.repo.QueryOneCart(user, cartId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return cart, nil
}

func (s *Service) CreateAddress(user *api.User, addr *api.Address) error {
	if user == nil || addr == nil {
		return errors.New("user or addr is nil")
	}
	//校验国家
	if len(addr.Country) <= 0 {
		return errors.New("country is empty")
	}
	if len(addr.Province) <= 0 {
		return errors.New("province is empty")
	}
	if len(addr.City) <= 0 {
		return errors.New("city is empty")
	}
	if len(addr.Phone) <= 0 {
		return errors.New("phone is empty")
	}
	//获取该用户已有地址数量
	_, total, err := s.repo.QueryAddress(user)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	addr.Id = total + 1
	addr.CreateTime = times.Timestamp()
	addr.UpdateTime = times.Timestamp()
	addr.DeleteTime = 0
	addr.IsDefault  = false

	err = s.repo.CreateAddress(user, addr)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) QueryAddress(user *api.User) ([]api.Address, int, error) {
	if user == nil {
		return nil, 0, errors.New("user is nil")
	}
	address, total, err := s.repo.QueryAddress(user)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return address, int(total), nil
}

func (s *Service) CreateCreditCard(user *api.User, credit *api.CreditCard) error {
	if user == nil || credit == nil {
		return errors.New("user or credit is nil")
	}
	if len(credit.Bin) <= 0 {
		return errors.New("bin is empty")
	}
	if len(credit.Number) <= 0 {
		return errors.New("number is empty")
	}
	if len(credit.Expire) <= 0 {
		return errors.New("expire is empty")
	}
	if len(credit.CVV) <= 0 {
		return errors.New("cvv is empty")
	}
	//获取该用户已有信用卡数量
	_, total, err := s.repo.QueryCreditCard(user)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	credit.Id = total + 1
	credit.CreateTime = times.Timestamp()
	credit.UpdateTime = times.Timestamp()
	credit.DeleteTime = 0
	credit.IsDefault  = false

	err = s.repo.CreateCreditCard(user, credit)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *Service) QueryCreditCard(user *api.User) ([]api.CreditCard, int, error) {
	if user == nil {
		return nil, 0, errors.New("user is nil")
	}
	credits, total, err := s.repo.QueryCreditCard(user)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	//脱敏
	for i := range credits {
		length := len(credits[i].Number)
		credits[i].Number = fmt.Sprintf("%s%s%s", credits[i].Number[:4], "******", credits[i].Number[length-4:length])
		credits[i].CVV    = ""
	}
	return credits, int(total), nil
}