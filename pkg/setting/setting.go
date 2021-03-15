package setting

import "github.com/spf13/viper"

type Setting struct{
	vp *viper.Viper
}
func NewString() (*Setting,error){
	vp := viper.New()
	vp.SetConfigName("configs")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil{
		return nil,err
	}
	return &Setting{vp},nil
}
