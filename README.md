# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [search.proto](#search-proto)
    - [GetSocietyIn](#-GetSocietyIn)
    - [GetSocietyOut](#-GetSocietyOut)
    - [GetUserWithLimitIn](#-GetUserWithLimitIn)
    - [GetUserWithLimitOut](#-GetUserWithLimitOut)
    - [Society](#-Society)
    - [UserSr](#-UserSr)
  
    - [SearchService](#-SearchService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="search-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## search.proto



<a name="-GetSocietyIn"></a>

### GetSocietyIn
Данные запроса собществ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| part_name | [string](#string) |  | полное имя сообщества или часть имени, по которому будет осуществлен поиск |
| limit | [int64](#int64) |  |  |
| offset | [int64](#int64) |  |  |






<a name="-GetSocietyOut"></a>

### GetSocietyOut
Объект ответа на запрос поиска


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societies | [Society](#Society) | repeated | Массив объектов сообществ, удовлетворяющих поску |
| total | [int64](#int64) |  | Общее количество сообществ |






<a name="-GetUserWithLimitIn"></a>

### GetUserWithLimitIn
Объект запроса пользователей


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | количество запрашиваемых записей |
| offset | [int64](#int64) |  | offset записей |
| nickname | [string](#string) |  | nickname пользователя которого ищем |






<a name="-GetUserWithLimitOut"></a>

### GetUserWithLimitOut
Объект ответа на запрос пользователей


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserSr](#UserSr) | repeated | Срез пользователей |
| total | [int64](#int64) |  | Количество возвращаемых пользователей |






<a name="-Society"></a>

### Society
Объект описывающий сообщество


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Имя сообщества |
| avatar_link | [string](#string) |  | Аватарка сообщества |
| is_private | [bool](#bool) |  | Признак публичности сообщества |
| description | [string](#string) |  | Описание сообщества |






<a name="-UserSr"></a>

### UserSr
Объект пользователь


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) |  | Ник пользователя |
| uuid | [string](#string) |  | UUID пользователя |
| avatar_link | [string](#string) |  | ссылка на последний аватар пользователя |
| name | [string](#string) |  | имя пользователя |
| surname | [string](#string) |  | фамилия пользователя |





 

 

 


<a name="-SearchService"></a>

### SearchService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSociety | [.GetSocietyIn](#GetSocietyIn) | [.GetSocietyOut](#GetSocietyOut) | Метод получения сообществ в поиске |
| GetUserWithLimit | [.GetUserWithLimitIn](#GetUserWithLimitIn) | [.GetUserWithLimitOut](#GetUserWithLimitOut) | Метод получения списка users |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

