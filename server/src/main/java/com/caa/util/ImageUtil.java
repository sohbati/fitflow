package com.caa.util;

import com.caa.config.ConfigData;
import com.caa.config.ProjectConfig;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.multipart.MultipartFile;

import javax.imageio.ImageIO;
import java.awt.*;
import java.awt.image.BufferedImage;
import java.io.*;
import java.nio.file.Files;
import java.util.Base64;

/**
 * Created by Reza on 22/08/2018.
 */
public class ImageUtil {

    public static void saveToFileSystem(
            String folderName,
            String fileName,
            String type,
            byte[] imageByteArray)
            throws IOException {

        String confFolder = PublicUtil.getProjectConfigFolder();
        String folder = confFolder + "/" + folderName;
        String fullPath = folder + "/" + fileName + "." + type;
        File folderObj = new File(folder);
        if (!folderObj.exists()) {
            folderObj.mkdir();
        }

        FileOutputStream fos = new FileOutputStream(fullPath);
        try {
            fos.write(imageByteArray);
        }
        finally {
            fos.close();
        }
    }

    public static String encodeImage(byte[] image) {
        if (image == null || image.length == 0) {
            return "";
        }
        return Base64.getEncoder().withoutPadding().encodeToString(image);
    }

    public static String encodeImage(ByteArrayOutputStream stream) {
        if (stream == null) {
            return  "";
        }
        return Base64.getEncoder().withoutPadding().encodeToString(stream.toByteArray());
    }

    public static byte[] handleFileUpload(MultipartFile picture) {
        byte[] bytes = new byte[0];
        if (!picture.isEmpty()) {
            try {
                bytes = picture.getBytes();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
        return bytes;
    }

    public static byte[] shrinkImage(byte[] imageInByte, String imageSuffix) throws IOException {
        if (imageInByte.length == 0) {
            return new byte[0];
        }

        int w = PublicUtil.getPublicConfig().getShrinkedImageWidth();
        int h = PublicUtil.getPublicConfig().getShrinkedImageHeight();

        BufferedImage originalImage = toBufferedImage(imageInByte);
        int type = originalImage.getType() == 0? BufferedImage.TYPE_INT_ARGB : originalImage.getType();

        BufferedImage resizedImage = new BufferedImage(w, h, type);
        Graphics2D g = resizedImage.createGraphics();
        g.drawImage(originalImage, 0, 0, w, h, null);

        g.setComposite(AlphaComposite.Src);

        g.setRenderingHint(RenderingHints.KEY_INTERPOLATION,
                RenderingHints.VALUE_INTERPOLATION_BILINEAR);
        g.setRenderingHint(RenderingHints.KEY_RENDERING,
                RenderingHints.VALUE_RENDER_QUALITY);
        g.setRenderingHint(RenderingHints.KEY_ANTIALIASING,
                RenderingHints.VALUE_ANTIALIAS_ON);
        g.dispose();

        return toByte(resizedImage, imageSuffix);

    }

    private static byte[] toByte(BufferedImage bufferedImage, String imageFormat) throws IOException{
        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        ImageIO.write( bufferedImage, imageFormat, baos );
        byte[] imageInByte = baos.toByteArray();
        baos.flush();
        baos.close();
        return imageInByte;
    }

    private static BufferedImage toBufferedImage(byte[] imageInByte) throws IOException{
        InputStream in = new ByteArrayInputStream(imageInByte);
        BufferedImage img = ImageIO.read(in);
        return img;
    }

    public static String getImageFileSuffixTypeByContentType(String contentType) throws IOException {
        switch (contentType) {
            case "image/png" : return "png";
            case "image/jpg" : return "jpg";
            case "image/gif" : return "gif";
        }
        return "jpg";
    }

    public static String loadImageInEncodedString(String relativePath, String fileName) throws IOException {
        String fileFullPath = PublicUtil.getProjectConfigFolder() + "/" +
                relativePath + "/" + fileName;
        File f = new File(fileFullPath);
        if (!f.exists()) {
            return "";
        }
        byte[] bytes = Files.readAllBytes(f.toPath());
        return ImageUtil.encodeImage(bytes);
    }

    public static byte[] loadImageInbytes(String relativePath, String fileName) throws IOException {
        String fileFullPath = PublicUtil.getProjectConfigFolder() + "/" +
                relativePath + "/" + fileName;
        File f = new File(fileFullPath);
        if (!f.exists()) {
            return new byte[0];
        }
        return  Files.readAllBytes(f.toPath());
    }
}
