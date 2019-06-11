// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20181119

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ArithmeticOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *ArithmeticOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArithmeticOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ArithmeticOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽���ı���Ϣ��������������������ӡ�
		TextDetections []*TextArithmetic `json:"TextDetections,omitempty" name:"TextDetections" list`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArithmeticOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArithmeticOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Coord struct {

	// ������
	X *int64 `json:"X,omitempty" name:"X"`

	// ������
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type EnglishOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *EnglishOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnglishOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnglishOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽���ı���Ϣ��������������������ӡ�
		TextDetections []*TextDetectionEn `json:"TextDetections,omitempty" name:"TextDetections" list`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnglishOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnglishOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralAccurateOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *GeneralAccurateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralAccurateOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽���ı���Ϣ��������������������ӡ�
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralAccurateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralAccurateOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// �����ֶΡ�
	Scene *string `json:"Scene,omitempty" name:"Scene"`
}

func (r *GeneralBasicOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽���ı���Ϣ��������������������ӡ�
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// ��⵽�����ԣ�Ŀǰ֧�ֵ����ַ�ΧΪ���������ġ��������ġ�Ӣ�ġ����ġ����ġ�δ����½�������Ը������ֵ�֧�֡�
	// ���ؽ������Ϊ��zh-��Ӣ��ϣ�jap-���ģ�kor-���ġ�
		Language *string `json:"Language,omitempty" name:"Language"`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralBasicOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *GeneralFastOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralFastOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽���ı���Ϣ��������������������ӡ�
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// ��⵽�����ԣ�Ŀǰ֧�ֵ����ַ�ΧΪ���������ġ��������ġ�Ӣ�ġ����ġ����ġ�δ����½�������Ը������ֵ�֧�֡�
	// ���ؽ������Ϊ��zh - ��Ӣ��ϣ�jap - ���ģ�kor - ���ġ�
		Language *string `json:"Language,omitempty" name:"Language"`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralFastOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralFastOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IDCardOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// FRONT Ϊ���֤����Ƭ��һ�棨�����棩��
	// BACK Ϊ���֤�й��յ�һ�棨�����棩��
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`

	// ��ѡ�ֶΣ�������Ҫѡ���Ƿ������Ӧ�ֶΡ�
	// Ŀǰ�������ֶ�Ϊ��
	// CropIdCard�����֤��Ƭ�ü���bool ���ͣ�
	// CropPortrait��������Ƭ�ü���bool ���ͣ�
	// CopyWarn����ӡ���澯��bool ���ͣ�
	// ReshootWarn�����ĸ澯��bool ���͡�
	// 
	// SDK ���÷�ʽ�ο���
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// API 3.0 Explorer ���÷�ʽ�ο���
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *IDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IDCardOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// �����������棩
		Name *string `json:"Name,omitempty" name:"Name"`

		// �Ա������棩
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// ���壨�����棩
		Nation *string `json:"Nation,omitempty" name:"Nation"`

		// �������ڣ������棩
		Birth *string `json:"Birth,omitempty" name:"Birth"`

		// ��ַ�������棩
		Address *string `json:"Address,omitempty" name:"Address"`

		// ���֤�ţ������棩
		IdNum *string `json:"IdNum,omitempty" name:"IdNum"`

		// ��֤���أ������棩
		Authority *string `json:"Authority,omitempty" name:"Authority"`

		// ֤����Ч�ڣ������棩
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// ��չ��Ϣ����������Ŀ�ѡ�ֶη��ض�Ӧ���ݣ��������򲻷��أ���������ο�ʾ��3��Ŀǰ֧�ֵ���չ�ֶ�Ϊ��
	// IdCard�����֤��Ƭ������ CropIdCard ʱ���أ�
	// Portrait��������Ƭ������ CropPortrait ʱ���أ�
	// WarnInfos���澯��Ϣ��Code - �澯�룬Msg - �澯��Ϣ���ݣ���ʶ������ļ���ӡ��ʱ���ء�
	// 
	// Code �澯���б�����壺
	// -9103	���֤���ĸ澯��
	// -9102	���֤��ӡ���澯��
		AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IDCardOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TableOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *TableOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TableOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TableOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽���ı���Ϣ��������������������ӡ�
		TextDetections []*TextTable `json:"TextDetections,omitempty" name:"TextDetections" list`

		// Base64 ������ Excel ���ݡ�
		Data *string `json:"Data,omitempty" name:"Data"`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TableOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TableOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextArithmetic struct {

	// ʶ������ı�������
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// ���
	Result *bool `json:"Result,omitempty" name:"Result"`

	// ���Ŷ� 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// �ı������꣬���ĸ����������ʾ
	// ע�⣺���ֶο��ܷ��� null����ʾȡ������Чֵ��
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// ���ֶ�Ϊ��չ�ֶΡ�
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextDetection struct {

	// ʶ������ı�������
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// ���Ŷ� 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// �ı������꣬���ĸ����������ʾ
	// ע�⣺���ֶο��ܷ��� null����ʾȡ������Чֵ��
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// ���ֶ�Ϊ��չ�ֶΡ�
	// GeneralBasicOcr�ӿڷ��ض�����ϢParag������ParagNo��
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextDetectionEn struct {

	// ʶ������ı�������
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// ���Ŷ� 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// �ı������꣬���ĸ����������ʾ
	// ע�⣺���ֶο��ܷ��� null����ʾȡ������Чֵ��
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// ���ֶ�Ϊ��չ�ֶΡ�ĿǰEnglishOCR�ӿڷ�������Ϊ�ա�
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextTable struct {

	// ��Ԫ�����Ͻǵ�������
	ColTl *int64 `json:"ColTl,omitempty" name:"ColTl"`

	// ��Ԫ�����Ͻǵ�������
	RowTl *int64 `json:"RowTl,omitempty" name:"RowTl"`

	// ��Ԫ�����½ǵ�������
	ColBr *int64 `json:"ColBr,omitempty" name:"ColBr"`

	// ��Ԫ�����½ǵ�������
	RowBr *int64 `json:"RowBr,omitempty" name:"RowBr"`

	// ��Ԫ������
	Text *string `json:"Text,omitempty" name:"Text"`

	// ��Ԫ�����ͣ�����body��������壩��header����ͷ����footer����β������
	Type *string `json:"Type,omitempty" name:"Type"`

	// ���Ŷ� 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// �ı������꣬���ĸ����������ʾ
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// ���ֶ�Ϊ��չ�ֶ�
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextWaybill struct {

	// �ռ�������
	RecName *WaybillObj `json:"RecName,omitempty" name:"RecName"`

	// �ռ����ֻ���
	RecNum *WaybillObj `json:"RecNum,omitempty" name:"RecNum"`

	// �ռ��˵�ַ
	RecAddr *WaybillObj `json:"RecAddr,omitempty" name:"RecAddr"`

	// �ļ�������
	SenderName *WaybillObj `json:"SenderName,omitempty" name:"SenderName"`

	// �ļ����ֻ���
	SenderNum *WaybillObj `json:"SenderNum,omitempty" name:"SenderNum"`

	// �ļ��˵�ַ
	SenderAddr *WaybillObj `json:"SenderAddr,omitempty" name:"SenderAddr"`

	// �˵���
	WaybillNum *WaybillObj `json:"WaybillNum,omitempty" name:"WaybillNum"`
}

type VinOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *VinOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VinOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VinOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽�ĳ��� VIN �롣
		Vin *string `json:"Vin,omitempty" name:"Vin"`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VinOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VinOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WaybillOCRRequest struct {
	*tchttp.BaseRequest

	// ͼƬ�� Base64 ֵ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ��Base64����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�� ImageUrl��ImageBase64 �����ṩһ����������ṩ��ֻʹ�� ImageUrl��
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// ͼƬ�� Url ��ַ��
	// ֧�ֵ�ͼƬ��ʽ��PNG��JPG��JPEG���ݲ�֧�� GIF ��ʽ��
	// ֧�ֵ�ͼƬ��С��������ͼƬ�� Base64 ����󲻳��� 3M��ͼƬ����ʱ�䲻���� 3 �롣
	// ͼƬ�洢����Ѷ�Ƶ� Url �ɱ��ϸ��ߵ������ٶȺ��ȶ��ԣ�����ͼƬ�洢����Ѷ�ơ�
	// ����Ѷ�ƴ洢�� Url �ٶȺ��ȶ��Կ�����һ��Ӱ�졣
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *WaybillOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WaybillOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WaybillOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ��⵽���ı���Ϣ��������������������ӡ�
	// ע�⣺���ֶο��ܷ��� null����ʾȡ������Чֵ��
		TextDetections *TextWaybill `json:"TextDetections,omitempty" name:"TextDetections"`

		// Ψһ���� ID��ÿ�����󶼻᷵�ء���λ����ʱ��Ҫ�ṩ�ô������ RequestId��
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WaybillOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WaybillOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WaybillObj struct {

	// ʶ������ı�������
	Text *string `json:"Text,omitempty" name:"Text"`
}
