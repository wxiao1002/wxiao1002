package com.github.xiao;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.ResponseBody;

import javax.net.ssl.*;
import java.io.*;
import java.security.GeneralSecurityException;
import java.security.KeyStore;
import java.security.SecureRandom;
import java.security.cert.Certificate;
import java.security.cert.CertificateFactory;
import java.util.Arrays;
import java.util.Collection;
import java.util.concurrent.TimeUnit;

public class OKHTTP {
    private static OkHttpClient client;
    private static X509TrustManager trustManager;
    static
    {
        try
        {
            //这里是服务器的证书文件,笔者查看了其他的教程,使用的是getAssets().open(),那是AS的工程
            //这里是Maven工程,证书文件放在了src/main/resources下
            //可以以pem或crt结尾,具体可以向购买证书的服务商查询.
            trustManager = trustManagerForCertificates(new FileInputStream("src/main/resources/server.crt"));
            client = new OkHttpClient.Builder()
                .sslSocketFactory(createSSLSocketFactory(), trustManager)
                .hostnameVerifier((hostname, sslSession) -> {
                    //验证主机名
                    if("www.test.com".equals(hostname))
                    {
                        return true;
                    }
                    else
                    {
                        HostnameVerifier verifier = HttpsURLConnection.getDefaultHostnameVerifier();
                            return verifier.verify(hostname,sslSession);
                    }
                }).build();
        }
        catch (GeneralSecurityException | FileNotFoundException e)
        {
            e.printStackTrace();
        }
    }

    public static String send(String url)
    {
        Request request = new Request.Builder().url(url).build();
        //如果想要加上get/post请求的话再.build前添加即可   
        try (Response response = client.newCall(request).execute())
        {
            ResponseBody body = response.body();
            return body == null ? null : body.string();
        }
        catch (IOException e)
        {
            e.printStackTrace();
            return null;
        }
    }
    
    //以下代码为别人的轮子
    private static X509TrustManager trustManagerForCertificates(InputStream in)
            throws GeneralSecurityException
    {
        CertificateFactory certificateFactory = CertificateFactory.getInstance("X.509");
        Collection<? extends Certificate> certificates = certificateFactory.generateCertificates(in);
        if (certificates.isEmpty()) {
            throw new IllegalArgumentException("expected non-empty set of trusted certificates");
        }

        char[] password = "password".toCharArray(); // 这里可以使用任意密码
        KeyStore keyStore = newEmptyKeyStore(password);
        int index = 0;
        for (Certificate certificate : certificates) {
            String certificateAlias = Integer.toString(index++);
            keyStore.setCertificateEntry(certificateAlias, certificate);
        }

        // Use it to build an X509 trust manager.
        KeyManagerFactory keyManagerFactory = KeyManagerFactory.getInstance(
                KeyManagerFactory.getDefaultAlgorithm());
        keyManagerFactory.init(keyStore, password);
        TrustManagerFactory trustManagerFactory = TrustManagerFactory.getInstance(
                TrustManagerFactory.getDefaultAlgorithm());
        trustManagerFactory.init(keyStore);
        TrustManager[] trustManagers = trustManagerFactory.getTrustManagers();
        if (trustManagers.length != 1 || !(trustManagers[0] instanceof X509TrustManager))
        {
            throw new IllegalStateException("Unexpected default trust managers:" + Arrays.toString(trustManagers));
        }
        return (X509TrustManager) trustManagers[0];
    }

    private static KeyStore newEmptyKeyStore(char[] password) throws GeneralSecurityException {
        try {
            KeyStore keyStore = KeyStore.getInstance(KeyStore.getDefaultType()); // 这里添加自定义的密码，默认
            InputStream in = null; // By convention, 'null' creates an empty key store.
            keyStore.load(in, password);
            return keyStore;
        } catch (IOException e) {
            throw new AssertionError(e);
        }
    }

    private static SSLSocketFactory createSSLSocketFactory() {
        SSLSocketFactory ssfFactory = null;
        try {
            SSLContext sc = SSLContext.getInstance("TLS");
            sc.init(null, new TrustManager[]{trustManager}, new SecureRandom());
            ssfFactory = sc.getSocketFactory();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return ssfFactory;
    }
}
