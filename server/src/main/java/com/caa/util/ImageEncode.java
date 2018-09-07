package com.caa.util;

import java.io.ByteArrayOutputStream;
import java.util.Base64;

/**
 * Created by Reza on 22/08/2018.
 */
public class ImageEncode {

    public static String encodeImage(byte[] image) {
        return Base64.getEncoder().withoutPadding().encodeToString(image);
    }

    public static String encodeImage(ByteArrayOutputStream stream) {
        return Base64.getEncoder().withoutPadding().encodeToString(stream.toByteArray());
    }
}
