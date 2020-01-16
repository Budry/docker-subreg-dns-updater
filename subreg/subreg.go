package subreg

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Login_Container struct {
	Response *Login_Response `xml:"response,omitempty"`
}

type Login struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Login"`

	Login string `xml:"login,omitempty"`

	Password string `xml:"password,omitempty"`
}

type Login_Response struct {
	Status string `xml:"status,omitempty"`

	Data *Login_Data `xml:"data,omitempty"`

	Error *Error_Info `xml:"error,omitempty"`
}

type Login_Data struct {
	Ssid string `xml:"ssid,omitempty"`
}

type Error_Info struct {
	Errormsg string `xml:"errormsg,omitempty"`

	Errorcode *Error_Codes `xml:"errorcode,omitempty"`
}

type Error_Codes struct {
	Major int32 `xml:"major,omitempty"`

	Minor int32 `xml:"minor,omitempty"`
}


type Get_DNS_Zone_Container struct {
	Response *Get_DNS_Zone_Response `xml:"response,omitempty"`
}

type Get_DNS_Zone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Zone"`

	Ssid string `xml:"ssid,omitempty"`

	Domain string `xml:"domain,omitempty"`
}

type Get_DNS_Zone_Response struct {
	Status string `xml:"status,omitempty"`

	Data *Get_DNS_Zone_Data `xml:"data,omitempty"`

	Error *Error_Info `xml:"error,omitempty"`
}

type Get_DNS_Zone_Data struct {
	Domain string `xml:"domain,omitempty"`

	Records []*Get_DNS_Zone_Record `xml:"records,omitempty"`
}

type Get_DNS_Zone_Record struct {
	Id int32 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Content string `xml:"content,omitempty"`

	Prio int32 `xml:"prio,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

//type Check_Domain_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Container"`
//
//	Response *Check_Domain_Response `xml:"response,omitempty"`
//}
//
//type Check_Domain struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Params *Check_Domain_Params `xml:"params,omitempty"`
//}
//
//type Info_Domain_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Container"`
//
//	Response *Info_Domain_Response `xml:"response,omitempty"`
//}
//
//type Info_Domain struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//}
//
//type Info_Domain_CZ_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Container"`
//
//	Response *Info_Domain_CZ_Response `xml:"response,omitempty"`
//}
//
//type Info_Domain_CZ struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//}
//
//type Domains_List_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Container"`
//
//	Response *Domains_List_Response `xml:"response,omitempty"`
//}
//
//type Domains_List struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type Set_Autorenew_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew_Container"`
//
//	Response *Set_Autorenew_Response `xml:"response,omitempty"`
//}
//
//type Set_Autorenew struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Autorenew string `xml:"autorenew,omitempty"`
//}
//
//type Create_Contact_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Container"`
//
//	Response *Create_Contact_Response `xml:"response,omitempty"`
//}
//
//type Create_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Contact *Create_Contact_Contact `xml:"contact,omitempty"`
//}
//
//type Update_Contact_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Container"`
//
//	Response *Update_Contact_Response `xml:"response,omitempty"`
//}
//
//type Update_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Contact *Update_Contact_Contact `xml:"contact,omitempty"`
//}
//
//type Info_Contact_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Container"`
//
//	Response *Info_Contact_Response `xml:"response,omitempty"`
//}
//
//type Info_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Contact *Info_Contact_Contact `xml:"contact,omitempty"`
//}
//
//type Contacts_List_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Container"`
//
//	Response *Contacts_List_Response `xml:"response,omitempty"`
//}
//
//type Contacts_List struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type Check_Object_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object_Container"`
//
//	Response *Check_Object_Response `xml:"response,omitempty"`
//}
//
//type Check_Object struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Object string `xml:"object,omitempty"`
//
//	Id string `xml:"id,omitempty"`
//}
//
//type Info_Object_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Container"`
//
//	Response *Info_Object_Response `xml:"response,omitempty"`
//}
//
//type Info_Object struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Object string `xml:"object,omitempty"`
//
//	Id string `xml:"id,omitempty"`
//}
//
//type Make_Order_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Container"`
//
//	Response *Make_Order_Response `xml:"response,omitempty"`
//}
//
//type Make_Order struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Order *Make_Order_Order `xml:"order,omitempty"`
//}
//
//type Info_Order_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Container"`
//
//	Response *Info_Order_Response `xml:"response,omitempty"`
//}
//
//type Info_Order struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Order int32 `xml:"order,omitempty"`
//}
//
//type Get_Credit_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Container"`
//
//	Response *Get_Credit_Response `xml:"response,omitempty"`
//}
//
//type Get_Credit struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type Get_Accountings_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Container"`
//
//	Response *Get_Accountings_Response `xml:"response,omitempty"`
//}
//
//type Get_Accountings struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	From string `xml:"from,omitempty"`
//
//	To string `xml:"to,omitempty"`
//}
//
//type Client_Payment_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment_Container"`
//
//	Response *Client_Payment_Response `xml:"response,omitempty"`
//}
//
//type Client_Payment struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Username string `xml:"username,omitempty"`
//
//	Amount float64 `xml:"amount,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//}
//
//type Order_Payment_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Order_Payment_Container"`
//
//	Response *Order_Payment_Response `xml:"response,omitempty"`
//}
//
//type Order_Payment struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Order_Payment"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Payment int32 `xml:"payment,omitempty"`
//
//	Amount float64 `xml:"amount,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//}
//
//type Credit_Correction_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction_Container"`
//
//	Response *Credit_Correction_Response `xml:"response,omitempty"`
//}
//
//type Credit_Correction struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Username string `xml:"username,omitempty"`
//
//	Amount float64 `xml:"amount,omitempty"`
//
//	Reason string `xml:"reason,omitempty"`
//}
//
//type Pricelist_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Container"`
//
//	Response *Pricelist_Response `xml:"response,omitempty"`
//}
//
//type Pricelist struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type Prices_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Container"`
//
//	Response *Prices_Response `xml:"response,omitempty"`
//}
//
//type Prices struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Prices"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Tld string `xml:"tld,omitempty"`
//}
//
//type Get_Pricelist_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Container"`
//
//	Response *Get_Pricelist_Response `xml:"response,omitempty"`
//}
//
//type Get_Pricelist struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Pricelist string `xml:"pricelist,omitempty"`
//}
//
//type Set_Prices_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Container"`
//
//	Response *Set_Prices_Response `xml:"response,omitempty"`
//}
//
//type Set_Prices struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Pricelist string `xml:"pricelist,omitempty"`
//
//	Tld string `xml:"tld,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//
//	Prices []*Set_Prices_Price `xml:"prices,omitempty"`
//}
//
//type Download_Document_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document_Container"`
//
//	Response *Download_Document_Response `xml:"response,omitempty"`
//}
//
//type Download_Document struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Id int32 `xml:"id,omitempty"`
//}
//
//type Upload_Document_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document_Container"`
//
//	Response *Upload_Document_Response `xml:"response,omitempty"`
//}
//
//type Upload_Document struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Document string `xml:"document,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Filetype string `xml:"filetype,omitempty"`
//}
//
//type List_Documents_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Container"`
//
//	Response *List_Documents_Response `xml:"response,omitempty"`
//}
//
//type List_Documents struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type Users_List_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_Container"`
//
//	Response *Users_List_Response `xml:"response,omitempty"`
//}
//
//type Users_List struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Users_List"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type Anycast_ADD_Zone_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone_Container"`
//
//	Response *Anycast_ADD_Zone_Response `xml:"response,omitempty"`
//}
//
//type Anycast_ADD_Zone struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Server int32 `xml:"server,omitempty"`
//}
//
//type Anycast_Remove_Zone_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone_Container"`
//
//	Response *Anycast_Remove_Zone_Response `xml:"response,omitempty"`
//}
//
//type Anycast_Remove_Zone struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Server int32 `xml:"server,omitempty"`
//}
//
//
//type Add_DNS_Zone_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone_Container"`
//
//	Response *Add_DNS_Zone_Response `xml:"response,omitempty"`
//}
//
//type Add_DNS_Zone struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Template string `xml:"template,omitempty"`
//}
//
//type Delete_DNS_Zone_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone_Container"`
//
//	Response *Delete_DNS_Zone_Response `xml:"response,omitempty"`
//}
//
//type Delete_DNS_Zone struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//}
//
//type Set_DNS_Zone_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Container"`
//
//	Response *Set_DNS_Zone_Response `xml:"response,omitempty"`
//}
//
//type Set_DNS_Zone struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Records []*Set_DNS_Zone_Record `xml:"records,omitempty"`
//}

type Add_DNS_Record_Container struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Record_Container"`

	Response *Add_DNS_Record_Response `xml:"response,omitempty"`
}

type Add_DNS_Record struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Record"`

	Ssid string `xml:"ssid,omitempty"`

	Domain string `xml:"domain,omitempty"`

	Record *Add_DNS_Record_Record `xml:"record,omitempty"`
}

type Modify_DNS_Record_Container struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Modify_DNS_Record_Container"`

	Response *Modify_DNS_Record_Response `xml:"response,omitempty"`
}

type Modify_DNS_Record struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Modify_DNS_Record"`

	Ssid string `xml:"ssid,omitempty"`

	Domain string `xml:"domain,omitempty"`

	Record *Modify_DNS_Record_Record `xml:"record,omitempty"`
}

//type Delete_DNS_Record_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Container"`
//
//	Response *Delete_DNS_Record_Response `xml:"response,omitempty"`
//}
//
//type Delete_DNS_Record struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Record *Delete_DNS_Record_Record `xml:"record,omitempty"`
//}
//
//type POLL_Get_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get_Container"`
//
//	Response *POLL_Get_Response `xml:"response,omitempty"`
//}
//
//type POLL_Get struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type POLL_Ack_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack_Container"`
//
//	Response *POLL_Ack_Response `xml:"response,omitempty"`
//}
//
//type POLL_Ack struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Id int32 `xml:"id,omitempty"`
//}
//
//type OIB_Search_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Container"`
//
//	Response *OIB_Search_Response `xml:"response,omitempty"`
//}
//
//type OIB_Search struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Oib string `xml:"oib,omitempty"`
//}
//
//type Get_Certificate_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate_Container"`
//
//	Response *Get_Certificate_Response `xml:"response,omitempty"`
//}
//
//type Get_Certificate struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Orderid int32 `xml:"orderid,omitempty"`
//}
//
//type Get_Redirects_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects_Container"`
//
//	Response *Get_Redirects_Response `xml:"response,omitempty"`
//}
//
//type Get_Redirects struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//}
//
//type In_Subreg_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg_Container"`
//
//	Response *In_Subreg_Response `xml:"response,omitempty"`
//}
//
//type In_Subreg struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//}
//
//type Sign_DNS_Zone_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone_Container"`
//
//	Response *Sign_DNS_Zone_Response `xml:"response,omitempty"`
//}
//
//type Sign_DNS_Zone struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//}
//
//type Unsign_DNS_Zone_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone_Container"`
//
//	Response *Unsign_DNS_Zone_Response `xml:"response,omitempty"`
//}
//
//type Unsign_DNS_Zone struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//}
//
//type Get_DNS_Info_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Container"`
//
//	Response *Get_DNS_Info_Response `xml:"response,omitempty"`
//}
//
//type Get_DNS_Info struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Dnstype string `xml:"dnstype,omitempty"`
//}
//
//type Special_Pricelist_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Container"`
//
//	Response *Special_Pricelist_Response `xml:"response,omitempty"`
//}
//
//type Special_Pricelist struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist"`
//
//	Ssid string `xml:"ssid,omitempty"`
//}
//
//type Get_TLD_Info_Container struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Container"`
//
//	Response *Get_TLD_Info_Response `xml:"response,omitempty"`
//}
//
//type Get_TLD_Info struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info"`
//
//	Ssid string `xml:"ssid,omitempty"`
//
//	Tld string `xml:"tld,omitempty"`
//}
//
//type Check_Domain_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Check_Domain_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Check_Domain_Params struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Params"`
//
//	Lang_info string `xml:"lang_info,omitempty"`
//}
//
//type Check_Domain_Price struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Price"`
//
//	Amount float64 `xml:"amount,omitempty"`
//
//	Amount_with_trustee float64 `xml:"amount_with_trustee,omitempty"`
//
//	Premium int32 `xml:"premium,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//}
//
//type Check_Domain_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Data"`
//
//	Name string `xml:"name,omitempty"`
//
//	Avail int32 `xml:"avail,omitempty"`
//
//	Existing_claim_id string `xml:"existing_claim_id,omitempty"`
//
//	Price *Check_Domain_Price `xml:"price,omitempty"`
//}
//
//type Info_Domain_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Info_Domain_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Info_Domain_Contacts struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Contacts"`
//
//	Admin []*Info_Domain_Contact `xml:"admin,omitempty"`
//
//	Tech []*Info_Domain_Contact `xml:"tech,omitempty"`
//
//	Bill []*Info_Domain_Contact `xml:"bill,omitempty"`
//}
//
//type Info_Domain_Dsdata struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Dsdata"`
//
//	Tag string `xml:"tag,omitempty"`
//
//	Alg string `xml:"alg,omitempty"`
//
//	Digest_type string `xml:"digest_type,omitempty"`
//
//	Digest string `xml:"digest,omitempty"`
//}
//
//type Info_Domain_Options struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Options"`
//
//	Nsset string `xml:"nsset,omitempty"`
//
//	Keyset string `xml:"keyset,omitempty"`
//
//	Dsdata []*Info_Domain_Dsdata `xml:"dsdata,omitempty"`
//
//	Keygroup string `xml:"keygroup,omitempty"`
//
//	Quarantined string `xml:"quarantined,omitempty"`
//}
//
//type Info_Domain_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Data"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Contacts *Info_Domain_Contacts `xml:"contacts,omitempty"`
//
//	Hosts []string `xml:"hosts,omitempty"`
//
//	Registrant *Info_Domain_Contact `xml:"registrant,omitempty"`
//
//	ExDate string `xml:"exDate,omitempty"`
//
//	CrDate string `xml:"crDate,omitempty"`
//
//	TrDate string `xml:"trDate,omitempty"`
//
//	UpDate string `xml:"upDate,omitempty"`
//
//	Authid string `xml:"authid,omitempty"`
//
//	Status []string `xml:"status,omitempty"`
//
//	Rgp []string `xml:"rgp,omitempty"`
//
//	Autorenew int32 `xml:"autorenew,omitempty"`
//
//	Premium int32 `xml:"premium,omitempty"`
//
//	Price float64 `xml:"price,omitempty"`
//
//	Whoisproxy int32 `xml:"whoisproxy,omitempty"`
//
//	Options *Info_Domain_Options `xml:"options,omitempty"`
//}
//
//type Info_Domain_CZ_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Info_Domain_CZ_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Info_Domain_CZ_Contacts struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Contacts"`
//
//	Admin []*Info_Domain_CZ_Contact `xml:"admin,omitempty"`
//
//	Tech []*Info_Domain_CZ_Contact `xml:"tech,omitempty"`
//
//	Bill []*Info_Domain_CZ_Contact `xml:"bill,omitempty"`
//}
//
//type Info_Domain_CZ_Dsdata struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Dsdata"`
//
//	Tag string `xml:"tag,omitempty"`
//
//	Alg string `xml:"alg,omitempty"`
//
//	Digest_type string `xml:"digest_type,omitempty"`
//
//	Digest string `xml:"digest,omitempty"`
//}
//
//type Info_Domain_CZ_Options struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Options"`
//
//	Nsset string `xml:"nsset,omitempty"`
//
//	Keyset string `xml:"keyset,omitempty"`
//
//	Dsdata *Info_Domain_CZ_Dsdata `xml:"dsdata,omitempty"`
//
//	Keygroup string `xml:"keygroup,omitempty"`
//
//	Quarantined string `xml:"quarantined,omitempty"`
//}
//
//type Info_Domain_CZ_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Data"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Contacts *Info_Domain_CZ_Contacts `xml:"contacts,omitempty"`
//
//	Hosts []string `xml:"hosts,omitempty"`
//
//	Registrant *Info_Domain_CZ_Contact `xml:"registrant,omitempty"`
//
//	ExDate string `xml:"exDate,omitempty"`
//
//	CrDate string `xml:"crDate,omitempty"`
//
//	TrDate string `xml:"trDate,omitempty"`
//
//	UpDate string `xml:"upDate,omitempty"`
//
//	Status []string `xml:"status,omitempty"`
//
//	Rgp []string `xml:"rgp,omitempty"`
//
//	Autorenew int32 `xml:"autorenew,omitempty"`
//
//	Whoisproxy int32 `xml:"whoisproxy,omitempty"`
//
//	Options *Info_Domain_CZ_Options `xml:"options,omitempty"`
//}
//
//type Domains_List_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Domains_List_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Domains_List_Domain struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Domain"`
//
//	Name string `xml:"name,omitempty"`
//
//	Expire string `xml:"expire,omitempty"`
//
//	Autorenew int32 `xml:"autorenew,omitempty"`
//}
//
//type Domains_List_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Data"`
//
//	Count int32 `xml:"count,omitempty"`
//
//	Domains []*Domains_List_Domain `xml:"domains,omitempty"`
//}
//
//type Set_Autorenew_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Set_Autorenew_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Set_Autorenew_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew_Data"`
//}
//
//type Create_Contact_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Create_Contact_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Create_Contact_Params struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Params"`
//
//	Regid string `xml:"regid,omitempty"`
//
//	Notify_email string `xml:"notify_email,omitempty"`
//
//	Vat string `xml:"vat,omitempty"`
//
//	Ident_type string `xml:"ident_type,omitempty"`
//
//	Ident_number string `xml:"ident_number,omitempty"`
//
//	Disclose []string `xml:"disclose,omitempty"`
//}
//
//type Create_Contact_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Contact"`
//
//	Name string `xml:"name,omitempty"`
//
//	Surname string `xml:"surname,omitempty"`
//
//	Org string `xml:"org,omitempty"`
//
//	Street string `xml:"street,omitempty"`
//
//	City string `xml:"city,omitempty"`
//
//	Pc string `xml:"pc,omitempty"`
//
//	Sp string `xml:"sp,omitempty"`
//
//	Cc string `xml:"cc,omitempty"`
//
//	Phone string `xml:"phone,omitempty"`
//
//	Fax string `xml:"fax,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//
//	Params *Create_Contact_Params `xml:"params,omitempty"`
//}
//
//type Create_Contact_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Data"`
//
//	Contactid string `xml:"contactid,omitempty"`
//}
//
//type Update_Contact_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Update_Contact_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Update_Contact_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Contact"`
//
//	Id string `xml:"id,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Surname string `xml:"surname,omitempty"`
//
//	Org string `xml:"org,omitempty"`
//
//	Street string `xml:"street,omitempty"`
//
//	City string `xml:"city,omitempty"`
//
//	Pc string `xml:"pc,omitempty"`
//
//	Sp string `xml:"sp,omitempty"`
//
//	Cc string `xml:"cc,omitempty"`
//
//	Phone string `xml:"phone,omitempty"`
//
//	Fax string `xml:"fax,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//}
//
//type Update_Contact_Order struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Order"`
//
//	Register string `xml:"register,omitempty"`
//
//	Orderid int32 `xml:"orderid,omitempty"`
//}
//
//type Update_Contact_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Data"`
//
//	Orders []*Update_Contact_Order `xml:"orders,omitempty"`
//}
//
//type Info_Contact_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Info_Contact_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Info_Contact_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Contact"`
//
//	Id string `xml:"id,omitempty"`
//}
//
//type Info_Contact_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Data"`
//
//	Id string `xml:"id,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Surname string `xml:"surname,omitempty"`
//
//	Org string `xml:"org,omitempty"`
//
//	Street string `xml:"street,omitempty"`
//
//	City string `xml:"city,omitempty"`
//
//	Pc string `xml:"pc,omitempty"`
//
//	Sp string `xml:"sp,omitempty"`
//
//	Cc string `xml:"cc,omitempty"`
//
//	Phone string `xml:"phone,omitempty"`
//
//	Fax string `xml:"fax,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//}
//
//type Contacts_List_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Contacts_List_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Contacts_List_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Contact"`
//
//	Name string `xml:"name,omitempty"`
//
//	Surname string `xml:"surname,omitempty"`
//
//	Org string `xml:"org,omitempty"`
//
//	Street string `xml:"street,omitempty"`
//
//	City string `xml:"city,omitempty"`
//
//	Pc string `xml:"pc,omitempty"`
//
//	Sp string `xml:"sp,omitempty"`
//
//	Cc string `xml:"cc,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//
//	Phone string `xml:"phone,omitempty"`
//
//	Fax string `xml:"fax,omitempty"`
//
//	Id string `xml:"id,omitempty"`
//}
//
//type Contacts_List_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Data"`
//
//	Contacts []*Contacts_List_Contact `xml:"contacts,omitempty"`
//
//	Count int32 `xml:"count,omitempty"`
//}
//
//type Check_Object_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Check_Object_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Check_Object_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object_Data"`
//
//	Id string `xml:"id,omitempty"`
//
//	Avail int32 `xml:"avail,omitempty"`
//}
//
//type Info_Object_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Info_Object_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Info_Object_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Contact"`
//
//	Name string `xml:"name,omitempty"`
//
//	Org string `xml:"org,omitempty"`
//
//	Street string `xml:"street,omitempty"`
//
//	City string `xml:"city,omitempty"`
//
//	Pc string `xml:"pc,omitempty"`
//
//	Sp string `xml:"sp,omitempty"`
//
//	Cc string `xml:"cc,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//
//	Phone string `xml:"phone,omitempty"`
//
//	Fax string `xml:"fax,omitempty"`
//
//	Vat string `xml:"vat,omitempty"`
//
//	Notify_email string `xml:"notify_email,omitempty"`
//
//	Ident_type string `xml:"ident_type,omitempty"`
//
//	Ident_number string `xml:"ident_number,omitempty"`
//
//	ClID string `xml:"clID,omitempty"`
//
//	Hidden []string `xml:"hidden,omitempty"`
//
//	Statuses []string `xml:"statuses,omitempty"`
//}
//
//type Info_Object_Ns struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Ns"`
//
//	Host string `xml:"host,omitempty"`
//
//	Ip string `xml:"ip,omitempty"`
//}
//
//type Info_Object_Nsset struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Nsset"`
//
//	Tech string `xml:"tech,omitempty"`
//
//	Ns []*Info_Object_Ns `xml:"ns,omitempty"`
//
//	ClID string `xml:"clID,omitempty"`
//}
//
//type Info_Object_Dnskey struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Dnskey"`
//
//	Flags string `xml:"flags,omitempty"`
//
//	Protocol string `xml:"protocol,omitempty"`
//
//	Alg string `xml:"alg,omitempty"`
//
//	PubKey string `xml:"pubKey,omitempty"`
//}
//
//type Info_Object_Keyset struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Keyset"`
//
//	Tech string `xml:"tech,omitempty"`
//
//	Dnskey []*Info_Object_Dnskey `xml:"dnskey,omitempty"`
//
//	ClID string `xml:"clID,omitempty"`
//}
//
//type Info_Object_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Data"`
//
//	Id string `xml:"id,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Contact *Info_Object_Contact `xml:"contact,omitempty"`
//
//	Nsset *Info_Object_Nsset `xml:"nsset,omitempty"`
//
//	Keyset *Info_Object_Keyset `xml:"keyset,omitempty"`
//}
//
//type Make_Order_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Make_Order_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Make_Order_Contacts struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Contacts"`
//
//	Admin *Make_Order_Contact `xml:"admin,omitempty"`
//
//	Tech *Make_Order_Contact `xml:"tech,omitempty"`
//
//	Billing *Make_Order_Contact `xml:"billing,omitempty"`
//}
//
//type Make_Order_Host struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Host"`
//
//	Hostname string `xml:"hostname,omitempty"`
//
//	Ipv4 string `xml:"ipv4,omitempty"`
//
//	Ipv6 string `xml:"ipv6,omitempty"`
//}
//
//type Make_Order_Ns struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Ns"`
//
//	Hosts []*Make_Order_Host `xml:"hosts,omitempty"`
//
//	Nsset string `xml:"nsset,omitempty"`
//}
//
//type Make_Order_New struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_New"`
//
//	Registrant *Make_Order_Contact `xml:"registrant,omitempty"`
//
//	Admin *Make_Order_Contact `xml:"admin,omitempty"`
//
//	Tech *Make_Order_Contact `xml:"tech,omitempty"`
//
//	Billing *Make_Order_Contact `xml:"billing,omitempty"`
//
//	Ns *Make_Order_Ns `xml:"ns,omitempty"`
//}
//
//type Make_Order_Dsdata struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Dsdata"`
//
//	Tag string `xml:"tag,omitempty"`
//
//	Alg string `xml:"alg,omitempty"`
//
//	Digest_type string `xml:"digest_type,omitempty"`
//
//	Digest string `xml:"digest,omitempty"`
//}
//
//type Make_Order_Param struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Param"`
//
//	Dsdata []*Make_Order_Dsdata `xml:"dsdata,omitempty"`
//
//	Param string `xml:"param,omitempty"`
//
//	Value string `xml:"value,omitempty"`
//}
//
//type Make_Order_Params struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Params"`
//
//	Period int32 `xml:"period,omitempty"`
//
//	Registrant *Make_Order_Contact `xml:"registrant,omitempty"`
//
//	Contacts *Make_Order_Contacts `xml:"contacts,omitempty"`
//
//	Ns *Make_Order_Ns `xml:"ns,omitempty"`
//
//	New *Make_Order_New `xml:"new,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Registry string `xml:"registry,omitempty"`
//
//	Authid string `xml:"authid,omitempty"`
//
//	Params []*Make_Order_Param `xml:"params,omitempty"`
//
//	Newowner string `xml:"newowner,omitempty"`
//
//	Reason string `xml:"reason,omitempty"`
//
//	Nicd string `xml:"nicd,omitempty"`
//
//	Password string `xml:"password,omitempty"`
//
//	Hostname string `xml:"hostname,omitempty"`
//
//	Ipv4 string `xml:"ipv4,omitempty"`
//
//	Ipv6 string `xml:"ipv6,omitempty"`
//
//	Dnstemp string `xml:"dnstemp,omitempty"`
//
//	Statuses []string `xml:"statuses,omitempty"`
//
//	Autorenew int32 `xml:"autorenew,omitempty"`
//}
//
//type Make_Order_Order struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Order"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Object string `xml:"object,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Params *Make_Order_Params `xml:"params,omitempty"`
//}
//
//type Make_Order_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Data"`
//
//	Orderid string `xml:"orderid,omitempty"`
//}
//
//type Info_Order_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Info_Order_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Info_Order_Order struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Order"`
//
//	Id int32 `xml:"id,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Status string `xml:"status,omitempty"`
//
//	Errorcode string `xml:"errorcode,omitempty"`
//
//	Lastupdate string `xml:"lastupdate,omitempty"`
//
//	Message string `xml:"message,omitempty"`
//
//	Payed string `xml:"payed,omitempty"`
//
//	Amount float64 `xml:"amount,omitempty"`
//}
//
//type Info_Order_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Data"`
//
//	Order *Info_Order_Order `xml:"order,omitempty"`
//}
//
//type Get_Credit_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Get_Credit_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Get_Credit_Credit struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Credit"`
//
//	Amount float64 `xml:"amount,omitempty"`
//
//	Threshold float64 `xml:"threshold,omitempty"`
//
//	Users float64 `xml:"users,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//}
//
//type Get_Credit_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Data"`
//
//	Credit *Get_Credit_Credit `xml:"credit,omitempty"`
//}
//
//type Get_Accountings_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Get_Accountings_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Get_Accountings_Accounting struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Accounting"`
//
//	Date string `xml:"date,omitempty"`
//
//	Text string `xml:"text,omitempty"`
//
//	Order int32 `xml:"order,omitempty"`
//
//	Sum float64 `xml:"sum,omitempty"`
//
//	Credit float64 `xml:"credit,omitempty"`
//}
//
//type Get_Accountings_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Data"`
//
//	Count int32 `xml:"count,omitempty"`
//
//	From string `xml:"from,omitempty"`
//
//	To string `xml:"to,omitempty"`
//
//	Accounting []*Get_Accountings_Accounting `xml:"accounting,omitempty"`
//}
//
//type Client_Payment_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Client_Payment_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Client_Payment_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment_Data"`
//}
//
//type Order_Payment_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Order_Payment_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Order_Payment_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Order_Payment_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Order_Payment_Data"`
//}
//
//type Credit_Correction_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Credit_Correction_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Credit_Correction_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction_Data"`
//}
//
//type Pricelist_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Pricelist_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Pricelist_Price struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Price"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Value float64 `xml:"value,omitempty"`
//}
//
//type Pricelist_Value struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Value"`
//
//	Value string `xml:"value,omitempty"`
//
//	Description string `xml:"description,omitempty"`
//}
//
//type Pricelist_Param struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Param"`
//
//	Param string `xml:"param,omitempty"`
//
//	Desc string `xml:"desc,omitempty"`
//
//	Required int32 `xml:"required,omitempty"`
//
//	Error_code int32 `xml:"error_code,omitempty"`
//
//	Values []*Pricelist_Value `xml:"values,omitempty"`
//}
//
//type Pricelist_Pricelist struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Pricelist"`
//
//	Tld string `xml:"tld,omitempty"`
//
//	Promo int32 `xml:"promo,omitempty"`
//
//	Promoexp string `xml:"promoexp,omitempty"`
//
//	Country string `xml:"country,omitempty"`
//
//	Continent string `xml:"continent,omitempty"`
//
//	Minyear int32 `xml:"minyear,omitempty"`
//
//	Maxyear int32 `xml:"maxyear,omitempty"`
//
//	Minyear_renew int32 `xml:"minyear_renew,omitempty"`
//
//	Maxyear_renew int32 `xml:"maxyear_renew,omitempty"`
//
//	Local_presence int32 `xml:"local_presence,omitempty"`
//
//	Prices []*Pricelist_Price `xml:"prices,omitempty"`
//
//	Statuses []string `xml:"statuses,omitempty"`
//
//	Params []*Pricelist_Param `xml:"params,omitempty"`
//}
//
//type Pricelist_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Data"`
//
//	Pricelist []*Pricelist_Pricelist `xml:"pricelist,omitempty"`
//}
//
//type Prices_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Prices_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Prices_Price struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Price"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Value float64 `xml:"value,omitempty"`
//}
//
//type Prices_Value struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Value"`
//
//	Value string `xml:"value,omitempty"`
//
//	Description string `xml:"description,omitempty"`
//}
//
//type Prices_Param struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Param"`
//
//	Param string `xml:"param,omitempty"`
//
//	Desc string `xml:"desc,omitempty"`
//
//	Required int32 `xml:"required,omitempty"`
//
//	Error_code int32 `xml:"error_code,omitempty"`
//
//	Values []*Prices_Value `xml:"values,omitempty"`
//}
//
//type Prices_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Data"`
//
//	Tld string `xml:"tld,omitempty"`
//
//	Country string `xml:"country,omitempty"`
//
//	Continent string `xml:"continent,omitempty"`
//
//	Minyear int32 `xml:"minyear,omitempty"`
//
//	Maxyear int32 `xml:"maxyear,omitempty"`
//
//	Local_presence int32 `xml:"local_presence,omitempty"`
//
//	Prices []*Prices_Price `xml:"prices,omitempty"`
//
//	Statuses []string `xml:"statuses,omitempty"`
//
//	Params []*Prices_Param `xml:"params,omitempty"`
//}
//
//type Get_Pricelist_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Get_Pricelist_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Get_Pricelist_Price struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Price"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Value float64 `xml:"value,omitempty"`
//}
//
//type Get_Pricelist_Pricelist struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Pricelist"`
//
//	Tld string `xml:"tld,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//
//	Prices []*Get_Pricelist_Price `xml:"prices,omitempty"`
//}
//
//type Get_Pricelist_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Data"`
//
//	Pricelist []*Get_Pricelist_Pricelist `xml:"pricelist,omitempty"`
//}
//
//type Set_Prices_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Set_Prices_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Set_Prices_Price struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Price"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Value float64 `xml:"value,omitempty"`
//}
//
//type Set_Prices_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Data"`
//}
//
//type Download_Document_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Download_Document_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Download_Document_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document_Data"`
//
//	Id int32 `xml:"id,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Filetype string `xml:"filetype,omitempty"`
//
//	Account string `xml:"account,omitempty"`
//
//	Document string `xml:"document,omitempty"`
//}
//
//type Upload_Document_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Upload_Document_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Upload_Document_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document_Data"`
//
//	Id int32 `xml:"id,omitempty"`
//}
//
//type List_Documents_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *List_Documents_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type List_Documents_Document struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Document"`
//
//	Id string `xml:"id,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Filetype string `xml:"filetype,omitempty"`
//
//	Account string `xml:"account,omitempty"`
//
//	Orderid int32 `xml:"orderid,omitempty"`
//}
//
//type List_Documents_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Data"`
//
//	Documents []*List_Documents_Document `xml:"documents,omitempty"`
//}
//
//type Users_List_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Users_List_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Users_List_User struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_User"`
//
//	Id int32 `xml:"id,omitempty"`
//
//	Username string `xml:"username,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Credit string `xml:"credit,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//
//	Billing_name string `xml:"billing_name,omitempty"`
//
//	Billing_street string `xml:"billing_street,omitempty"`
//
//	Billing_city string `xml:"billing_city,omitempty"`
//
//	Billing_pc string `xml:"billing_pc,omitempty"`
//
//	Billing_country string `xml:"billing_country,omitempty"`
//
//	Company_id string `xml:"company_id,omitempty"`
//
//	Company_vat string `xml:"company_vat,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//
//	Phone string `xml:"phone,omitempty"`
//
//	Last_login string `xml:"last_login,omitempty"`
//}
//
//type Users_List_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_Data"`
//
//	Count int32 `xml:"count,omitempty"`
//
//	Users []*Users_List_User `xml:"users,omitempty"`
//}
//
//type Anycast_ADD_Zone_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Anycast_ADD_Zone_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Anycast_ADD_Zone_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone_Data"`
//}
//
//type Anycast_Remove_Zone_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Anycast_Remove_Zone_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Anycast_Remove_Zone_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone_Data"`
//}
//

//
//type Get_DNS_Zone_Record struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Zone_Record"`
//
//	Id int32 `xml:"id,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Content string `xml:"content,omitempty"`
//
//	Prio int32 `xml:"prio,omitempty"`
//
//	Ttl int32 `xml:"ttl,omitempty"`
//}
//

//
//type Add_DNS_Zone_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Add_DNS_Zone_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Add_DNS_Zone_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone_Data"`
//}
//
//type Delete_DNS_Zone_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Delete_DNS_Zone_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Delete_DNS_Zone_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone_Data"`
//}
//
//type Set_DNS_Zone_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Set_DNS_Zone_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Set_DNS_Zone_Record struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Record"`
//
//	Name string `xml:"name,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Content string `xml:"content,omitempty"`
//
//	Prio int32 `xml:"prio,omitempty"`
//
//	Ttl int32 `xml:"ttl,omitempty"`
//}
//
//type Set_DNS_Zone_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Data"`
//}

type Add_DNS_Record_Response struct {
	Status string `xml:"status,omitempty"`

	Data *Add_DNS_Record_Data `xml:"data,omitempty"`

	Error *Error_Info `xml:"error,omitempty"`
}

type Add_DNS_Record_Record struct {
	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Content string `xml:"content,omitempty"`

	Prio int32 `xml:"prio,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type Add_DNS_Record_Data struct {
}

type Modify_DNS_Record_Response struct {
	Status string `xml:"status,omitempty"`

	Data *Modify_DNS_Record_Data `xml:"data,omitempty"`

	Error *Error_Info `xml:"error,omitempty"`
}

type Modify_DNS_Record_Record struct {
	Id int32 `xml:"id,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Content string `xml:"content,omitempty"`

	Prio int32 `xml:"prio,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type Modify_DNS_Record_Data struct {
}

//type Delete_DNS_Record_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Delete_DNS_Record_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Delete_DNS_Record_Record struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Record"`
//
//	Id int32 `xml:"id,omitempty"`
//}
//
//type Delete_DNS_Record_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Data"`
//}
//
//type POLL_Get_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *POLL_Get_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type POLL_Get_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get_Data"`
//
//	Count int32 `xml:"count,omitempty"`
//
//	Date string `xml:"date,omitempty"`
//
//	Id int32 `xml:"id,omitempty"`
//
//	Orderid int32 `xml:"orderid,omitempty"`
//
//	Orderstatus string `xml:"orderstatus,omitempty"`
//
//	Message string `xml:"message,omitempty"`
//
//	Errorcode string `xml:"errorcode,omitempty"`
//}
//
//type POLL_Ack_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *POLL_Ack_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type POLL_Ack_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack_Data"`
//}
//
//type OIB_Search_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *OIB_Search_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type OIB_Search_Domain struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Domain"`
//
//	Name string `xml:"name,omitempty"`
//
//	Type_ int32 `xml:"type,omitempty"`
//
//	Typedesc string `xml:"typedesc,omitempty"`
//}
//
//type OIB_Search_Type struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Type"`
//
//	Type_ int32 `xml:"type,omitempty"`
//
//	Typedesc string `xml:"typedesc,omitempty"`
//
//	Used int32 `xml:"used,omitempty"`
//
//	Maximum int32 `xml:"maximum,omitempty"`
//}
//
//type OIB_Search_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Data"`
//
//	Domains []*OIB_Search_Domain `xml:"domains,omitempty"`
//
//	Types []*OIB_Search_Type `xml:"types,omitempty"`
//}
//
//type Get_Certificate_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Get_Certificate_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Get_Certificate_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate_Data"`
//
//	Certificate string `xml:"certificate,omitempty"`
//
//	Expire string `xml:"expire,omitempty"`
//
//	Domain string `xml:"domain,omitempty"`
//
//	Type_ string `xml:"type,omitempty"`
//}
//
//type Get_Redirects_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Get_Redirects_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Get_Redirects_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects_Data"`
//
//	Web string `xml:"web,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//}
//
//type In_Subreg_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *In_Subreg_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type In_Subreg_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg_Data"`
//
//	Myaccount string `xml:"myaccount,omitempty"`
//}
//
//type Sign_DNS_Zone_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Sign_DNS_Zone_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Sign_DNS_Zone_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone_Data"`
//}
//
//type Unsign_DNS_Zone_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Unsign_DNS_Zone_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Unsign_DNS_Zone_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone_Data"`
//}
//
//type Get_DNS_Info_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Get_DNS_Info_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Get_DNS_Info_Anydata struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Anydata"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Data string `xml:"data,omitempty"`
//}
//
//type Get_DNS_Info_Dn struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Dn"`
//
//	Nameserver string `xml:"nameserver,omitempty"`
//
//	Anydata []*Get_DNS_Info_Anydata `xml:"anydata,omitempty"`
//
//	Nslist []string `xml:"nslist,omitempty"`
//
//	Soaid string `xml:"soaid,omitempty"`
//}
//
//type Get_DNS_Info_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Data"`
//
//	In_zone string `xml:"in_zone,omitempty"`
//
//	Dnssec string `xml:"dnssec,omitempty"`
//
//	Dns []*Get_DNS_Info_Dn `xml:"dns,omitempty"`
//}
//
//type Special_Pricelist_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Special_Pricelist_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Special_Pricelist_Price struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Price"`
//
//	Register float64 `xml:"register,omitempty"`
//
//	Renew float64 `xml:"renew,omitempty"`
//
//	Transfer float64 `xml:"transfer,omitempty"`
//}
//
//type Special_Pricelist_Pricelist struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Pricelist"`
//
//	Tld string `xml:"tld,omitempty"`
//
//	Currency string `xml:"currency,omitempty"`
//
//	Dateto string `xml:"dateto,omitempty"`
//
//	Prices []*Special_Pricelist_Price `xml:"prices,omitempty"`
//}
//
//type Special_Pricelist_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Data"`
//
//	Pricelist []*Special_Pricelist_Pricelist `xml:"pricelist,omitempty"`
//}
//
//type Get_TLD_Info_Response struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Response"`
//
//	Status string `xml:"status,omitempty"`
//
//	Data *Get_TLD_Info_Data `xml:"data,omitempty"`
//
//	Error *Error_Info `xml:"error,omitempty"`
//}
//
//type Get_TLD_Info_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Contact"`
//
//	Type_ string `xml:"type,omitempty"`
//
//	Cnt int32 `xml:"cnt,omitempty"`
//}
//
//type Get_TLD_Info_Option struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Option"`
//
//	Value string `xml:"value,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//}
//
//type Get_TLD_Info_Param struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Param"`
//
//	Param string `xml:"param,omitempty"`
//
//	Name string `xml:"name,omitempty"`
//
//	Desc string `xml:"desc,omitempty"`
//
//	Required string `xml:"required,omitempty"`
//
//	Options []*Get_TLD_Info_Option `xml:"options,omitempty"`
//}
//
//type Get_TLD_Info_Data struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Data"`
//
//	PeriodsCreate []string `xml:"periodsCreate,omitempty"`
//
//	PeriodsRenew []string `xml:"periodsRenew,omitempty"`
//
//	Transfer string `xml:"transfer,omitempty"`
//
//	Trade string `xml:"trade,omitempty"`
//
//	Idn string `xml:"idn,omitempty"`
//
//	Trustee string `xml:"trustee,omitempty"`
//
//	Ns string `xml:"ns,omitempty"`
//
//	Contacts []*Get_TLD_Info_Contact `xml:"contacts,omitempty"`
//
//	Params []*Get_TLD_Info_Param `xml:"params,omitempty"`
//}
//type Info_Domain_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Contact"`
//
//	Subregid string `xml:"subregid,omitempty"`
//
//	Registryid string `xml:"registryid,omitempty"`
//}
//
//type Info_Domain_CZ_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Contact"`
//
//	Subregid string `xml:"subregid,omitempty"`
//
//	Registryid string `xml:"registryid,omitempty"`
//}
//
//type Make_Order_Contact_New struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Contact_New"`
//
//	Name string `xml:"name,omitempty"`
//
//	Surname string `xml:"surname,omitempty"`
//
//	Org string `xml:"org,omitempty"`
//
//	Street string `xml:"street,omitempty"`
//
//	City string `xml:"city,omitempty"`
//
//	Pc string `xml:"pc,omitempty"`
//
//	Sp string `xml:"sp,omitempty"`
//
//	Cc string `xml:"cc,omitempty"`
//
//	Phone string `xml:"phone,omitempty"`
//
//	Fax string `xml:"fax,omitempty"`
//
//	Email string `xml:"email,omitempty"`
//}
//
//type Make_Order_Contact struct {
//	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Contact"`
//
//	Id string `xml:"id,omitempty"`
//
//	Regid string `xml:"regid,omitempty"`
//
//	New *Make_Order_Contact_New `xml:"new,omitempty"`
//}

type SubregCz struct {
	client *SOAPClient
}

func NewSubregCz(url string, tls bool, auth *BasicAuth) *SubregCz {
	if url == "" {
		url = "https://subreg.cz/soap/cmd.php?soap_format=1"
	}
	client := NewSOAPClient(url, tls, auth)

	return &SubregCz{
		client: client,
	}
}

func NewSubregCzWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SubregCz {
	if url == "" {
		url = "https://subreg.cz/soap/cmd.php?soap_format=1"
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &SubregCz{
		client: client,
	}
}

func (service *SubregCz) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *SubregCz) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *SubregCz) Login(request *Login) (*Login_Container, error) {
	response := new(Login_Container)
	err := service.client.Call("http://subreg.cz/wsdl#Login", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//func (service *SubregCz) Check_Domain(request *Check_Domain) (*Check_Domain_Container, error) {
//	response := new(Check_Domain_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Check_Domain", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Info_Domain(request *Info_Domain) (*Info_Domain_Container, error) {
//	response := new(Info_Domain_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Info_Domain", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Info_Domain_CZ(request *Info_Domain_CZ) (*Info_Domain_CZ_Container, error) {
//	response := new(Info_Domain_CZ_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Info_Domain_CZ", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Domains_List(request *Domains_List) (*Domains_List_Container, error) {
//	response := new(Domains_List_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Domains_List", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Set_Autorenew(request *Set_Autorenew) (*Set_Autorenew_Container, error) {
//	response := new(Set_Autorenew_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Set_Autorenew", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Create_Contact(request *Create_Contact) (*Create_Contact_Container, error) {
//	response := new(Create_Contact_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Create_Contact", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Update_Contact(request *Update_Contact) (*Update_Contact_Container, error) {
//	response := new(Update_Contact_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Update_Contact", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Info_Contact(request *Info_Contact) (*Info_Contact_Container, error) {
//	response := new(Info_Contact_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Info_Contact", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Contacts_List(request *Contacts_List) (*Contacts_List_Container, error) {
//	response := new(Contacts_List_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Contacts_List", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Check_Object(request *Check_Object) (*Check_Object_Container, error) {
//	response := new(Check_Object_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Check_Object", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Info_Object(request *Info_Object) (*Info_Object_Container, error) {
//	response := new(Info_Object_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Info_Object", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Make_Order(request *Make_Order) (*Make_Order_Container, error) {
//	response := new(Make_Order_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Make_Order", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Info_Order(request *Info_Order) (*Info_Order_Container, error) {
//	response := new(Info_Order_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Info_Order", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Get_Credit(request *Get_Credit) (*Get_Credit_Container, error) {
//	response := new(Get_Credit_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Get_Credit", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Get_Accountings(request *Get_Accountings) (*Get_Accountings_Container, error) {
//	response := new(Get_Accountings_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Get_Accountings", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Client_Payment(request *Client_Payment) (*Client_Payment_Container, error) {
//	response := new(Client_Payment_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Client_Payment", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Order_Payment(request *Order_Payment) (*Order_Payment_Container, error) {
//	response := new(Order_Payment_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Order_Payment", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Credit_Correction(request *Credit_Correction) (*Credit_Correction_Container, error) {
//	response := new(Credit_Correction_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Credit_Correction", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Pricelist(request *Pricelist) (*Pricelist_Container, error) {
//	response := new(Pricelist_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Pricelist", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Prices(request *Prices) (*Prices_Container, error) {
//	response := new(Prices_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Prices", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Get_Pricelist(request *Get_Pricelist) (*Get_Pricelist_Container, error) {
//	response := new(Get_Pricelist_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Get_Pricelist", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Set_Prices(request *Set_Prices) (*Set_Prices_Container, error) {
//	response := new(Set_Prices_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Set_Prices", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Download_Document(request *Download_Document) (*Download_Document_Container, error) {
//	response := new(Download_Document_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Download_Document", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Upload_Document(request *Upload_Document) (*Upload_Document_Container, error) {
//	response := new(Upload_Document_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Upload_Document", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) List_Documents(request *List_Documents) (*List_Documents_Container, error) {
//	response := new(List_Documents_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#List_Documents", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Users_List(request *Users_List) (*Users_List_Container, error) {
//	response := new(Users_List_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Users_List", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Anycast_ADD_Zone(request *Anycast_ADD_Zone) (*Anycast_ADD_Zone_Container, error) {
//	response := new(Anycast_ADD_Zone_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Anycast_ADD_Zone", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Anycast_Remove_Zone(request *Anycast_Remove_Zone) (*Anycast_Remove_Zone_Container, error) {
//	response := new(Anycast_Remove_Zone_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Anycast_Remove_Zone", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
func (service *SubregCz) Get_DNS_Zone(request *Get_DNS_Zone) (*Get_DNS_Zone_Container, error) {
	response := new(Get_DNS_Zone_Container)
	err := service.client.Call("http://subreg.cz/wsdl#Get_DNS_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
//
//func (service *SubregCz) Add_DNS_Zone(request *Add_DNS_Zone) (*Add_DNS_Zone_Container, error) {
//	response := new(Add_DNS_Zone_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Add_DNS_Zone", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Delete_DNS_Zone(request *Delete_DNS_Zone) (*Delete_DNS_Zone_Container, error) {
//	response := new(Delete_DNS_Zone_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Delete_DNS_Zone", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Set_DNS_Zone(request *Set_DNS_Zone) (*Set_DNS_Zone_Container, error) {
//	response := new(Set_DNS_Zone_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Set_DNS_Zone", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
func (service *SubregCz) Add_DNS_Record(request *Add_DNS_Record) (*Add_DNS_Record_Container, error) {
	response := new(Add_DNS_Record_Container)
	err := service.client.Call("http://subreg.cz/wsdl#Add_DNS_Record", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
//
func (service *SubregCz) Modify_DNS_Record(request *Modify_DNS_Record) (*Modify_DNS_Record_Container, error) {
	response := new(Modify_DNS_Record_Container)
	err := service.client.Call("http://subreg.cz/wsdl#Modify_DNS_Record", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
//
//func (service *SubregCz) Delete_DNS_Record(request *Delete_DNS_Record) (*Delete_DNS_Record_Container, error) {
//	response := new(Delete_DNS_Record_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Delete_DNS_Record", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) POLL_Get(request *POLL_Get) (*POLL_Get_Container, error) {
//	response := new(POLL_Get_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#POLL_Get", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) POLL_Ack(request *POLL_Ack) (*POLL_Ack_Container, error) {
//	response := new(POLL_Ack_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#POLL_Ack", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) OIB_Search(request *OIB_Search) (*OIB_Search_Container, error) {
//	response := new(OIB_Search_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#OIB_Search", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Get_Certificate(request *Get_Certificate) (*Get_Certificate_Container, error) {
//	response := new(Get_Certificate_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Get_Certificate", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Get_Redirects(request *Get_Redirects) (*Get_Redirects_Container, error) {
//	response := new(Get_Redirects_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Get_Redirects", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) In_Subreg(request *In_Subreg) (*In_Subreg_Container, error) {
//	response := new(In_Subreg_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#In_Subreg", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Sign_DNS_Zone(request *Sign_DNS_Zone) (*Sign_DNS_Zone_Container, error) {
//	response := new(Sign_DNS_Zone_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Sign_DNS_Zone", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Unsign_DNS_Zone(request *Unsign_DNS_Zone) (*Unsign_DNS_Zone_Container, error) {
//	response := new(Unsign_DNS_Zone_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Unsign_DNS_Zone", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Get_DNS_Info(request *Get_DNS_Info) (*Get_DNS_Info_Container, error) {
//	response := new(Get_DNS_Info_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Get_DNS_Info", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Special_Pricelist(request *Special_Pricelist) (*Special_Pricelist_Container, error) {
//	response := new(Special_Pricelist_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Special_Pricelist", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (service *SubregCz) Get_TLD_Info(request *Get_TLD_Info) (*Get_TLD_Info_Container, error) {
//	response := new(Get_TLD_Info_Container)
//	err := service.client.Call("http://subreg.cz/wsdl#Get_TLD_Info", request, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Items []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tlsCfg  *tls.Config
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, insecureSkipVerify bool, auth *BasicAuth) *SOAPClient {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	return NewSOAPClientWithTLSConfig(url, tlsCfg, auth)
}

func NewSOAPClientWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:    url,
		tlsCfg: tlsCfg,
		auth:   auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.tlsCfg,
		Dial:            dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
