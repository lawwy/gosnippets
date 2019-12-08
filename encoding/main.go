package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// package elkstack.utils;

// import org.apache.commons.codec.binary.Base64;
// import org.apache.commons.codec.digest.DigestUtils;

// import javax.crypto.Cipher;
// import javax.crypto.spec.SecretKeySpec;

// /**
//  * @author 李智慧
//  * @date 2017年8月1日
//  * @desc AES
//  */
// public class AES {

//     private static byte[] encrypt(byte[] text, byte[] key) throws Exception {
// 	SecretKeySpec aesKey = new SecretKeySpec(key, "AES");
// 	Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
// 	cipher.init(Cipher.ENCRYPT_MODE, aesKey);
// 	return cipher.doFinal(text);
//     }

//     private static byte[] decrypt(byte[] text, byte[] key) throws Exception {
// 	SecretKeySpec aesKey = new SecretKeySpec(key, "AES");

// 	Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
// 	cipher.init(Cipher.DECRYPT_MODE, aesKey);
// 	return cipher.doFinal(text);
//     }

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func decrypt(src []byte, key []byte) string {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return string(src)

}

//     /**
//      * @author 李智慧
//      * @date 2017年8月1日
//      * @desc 加密
//      * @param text 明文
//      * @param key 密钥
//      */
//     public static String encodeAES(String text, String key) throws Exception {
// 	byte[] keybBytes = DigestUtils.md5(key);
// 	byte[] passwdBytes = text.getBytes();
// 	byte[] aesBytyes = encrypt(passwdBytes, keybBytes);
// 	return new String(Base64.encodeBase64(aesBytyes));
//     }

//     /**
//      * @author 李智慧
//      * @date 2017年8月1日
//      * @desc 解密
//      * @param password 密文
//      * @param key 密钥
//      */
//     public static String deCodeAES(String password, String key) throws Exception {
// 	byte[] keybBytes = DigestUtils.md5(key);
// 	byte[] debase64Bytes = Base64.decodeBase64(password.getBytes());
// 	return new String(decrypt(debase64Bytes, keybBytes));
//     }

func decode(password string, key string) string {
	keyBytes := md5.New().Sum([]byte(key))
	// debase64Bytes := []byte{}
	debase64Bytes, _ := base64.StdEncoding.DecodeString(password)
	return decrypt(debase64Bytes, keyBytes)

}

// func md5V(str string) string  {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }
// /**
//  * @param args
//  */
// public static void main(String[] args) {
//     /*String encodeAES="";
//     try {
//       // encodeAES = encodeAES("omI0luCNtvhgqVL8YR0KJHrfwvCw","FE5Pj442Y3tiQuUt");
//        //System.out.println("密文"+encodeAES);
//        // FE5Pj442Y3tiQuUt
//        // 544Tkt42U8dczWt1BC6Q==
//         //544Tkt 42 U8dczWt1BC6Q==
//         String deCodeAES = deCodeAES("Qs963ZFcD2M18FpbT1kXHg==","74ceeb4a3ecb4f0788e9a28bae08a175");
//         //String deCodeAES = deCodeAES("PD8pT4KLoZb1TKI/fLgjzQ==","DFA0964D1185BB40");
//       // System.out.println(encodeAES);
//       System.out.println("+==========="+deCodeAES);
//     } catch (Exception e) {
//         // TODO Auto-generated catch block
//         e.printStackTrace();
//        System.out.println(encodeAES);
//     }*/
//     try{
// 		System.out.println(deCodeAES("qPYMBaQemLydcWVo20Oy5Q==","hHMogstx1gcJQLcu"));
// 	}catch (Exception e){
//     	e.printStackTrace();
// 	}

// }

// }

func main() {
	fmt.Println(decode("qPYMBaQemLydcWVo20Oy5Q==", "hHMogstx1gcJQLcu"))
}
