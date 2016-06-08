package main
import(
	"./slack"
	"fmt"
	"errors"
	"strconv"
	// "strings"
)
func roll(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	// a := explode_cmd(s.Text)
	z,e := roll_tokenizer(ct, s.Text)
	Message(ct.Websocket, s, implode(z," "))
	Message(ct.Websocket, s, fmt.Sprintf("%+v\n%+v", z,e))
}
func roll_not_care_space(a string, i *int){
	for a[*i]==' '{
		(*i)++
	}
}
func roll_get_int(a string, i *int)(string,error){
	roll_not_care_space(a,i)
	if len(a)<*i+1{
		return "", errors.New("EOT")
	}
	switch a[*i]{
		case '0','1','2','3','4','5','6','7','8','9':
			var c string
			for *i<len(a) && (a[*i]=='0'||a[*i]=='1'||a[*i]=='2'||a[*i]=='3'||a[*i]=='4'||a[*i]=='5'||a[*i]=='6'||a[*i]=='7'||a[*i]=='8'||a[*i]=='9') {
				c+=string(a[*i])
				(*i)++
			}
			return c,nil
		default:
			return "", errors.New("NOT INT")
	}
}
func roll_get_dice(a string, i *int)(string,error){
	roll_not_care_space(a,i)
	if len(a)<*i+1{
		return "", errors.New("EOT")
	}
	switch a[*i]{
		case 'd','D':
			*i++
			return "d",nil
		default:
			return "", errors.New("NOT DICE")
	}
}
func roll_tokenizer(ct *CT, a string)([]string,[]string){
	var b []string
	var t []string
	for i:=0;i<len(a);i++{
		z,err := roll_get_int(a,&i)
		if err==nil{
				z1,err1 := roll_get_dice(a,&i)
				if err1==nil{
					z2,err2 := roll_get_int(a,&i)
					if err2==nil{
						var tot string
						num,_:=strconv.Atoi(z)
						for i:=0;i<num;i++{
							max,_:=strconv.Atoi(z2)
							tot+="`"+strconv.Itoa(ct.Random.Intn(max)+1)+"` "
						}
						b = append(b,tot)
						t = append(t,"-")
					}else{
						b = append(b,z)
						t = append(t,"-")
						b = append(b,z1)
						t = append(t,"-")
						b = append(b,"-")
						t = append(t,err2.Error())
					}
				}else{
					b = append(b,z)
					t = append(t,"-")
					b = append(b,"-")
					t = append(t,err1.Error())
				}
		}else{
			b = append(b,"-")
			t = append(t,err.Error())
		}
	}
	return b,t
}