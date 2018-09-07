package com.caa.report;

import com.caa.model.ProgramExerciseItem;
import com.caa.modelview.ProgramExerciseItemView;
import com.caa.util.ImageEncode;
import net.sf.jasperreports.engine.*;
import net.sf.jasperreports.engine.data.JRBeanCollectionDataSource;
import net.sf.jasperreports.engine.design.JasperDesign;
import net.sf.jasperreports.engine.xml.JRXmlLoader;
import org.springframework.http.CacheControl;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.util.Base64Utils;
import sun.misc.BASE64Encoder;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.*;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Created by Reza on 20/08/2018.
 */
public class ExportReport {

    /**
     *
     * @param exerciseItems
     * @return BASE64 encoded string
     */
    public static String getProgramExerciseAsImage(List<ProgramExercisesReportDTO> exerciseItems) {
        ByteArrayOutputStream out = getImageStream(exerciseItems);
        String imageString = ImageEncode.encodeImage(out);
        return imageString;
    }

    public static byte[] getProgramExerciseAsImageInBytes(List<ProgramExercisesReportDTO> exerciseItems) {
        ByteArrayOutputStream out = getImageStream(exerciseItems);
        return out.toByteArray();
    }


    private static ByteArrayOutputStream getImageStream(List<ProgramExercisesReportDTO> exerciseItems) {
        try {
            if (exerciseItems.size() == 0) {
                return null;
            }
            InputStream inputStream = new FileInputStream(
                    new java.io.File(
                           exerciseItems.get(0).getClass().getClassLoader().getResource(
                                   "report/report1.jrxml").getFile()));

            JasperDesign jasperDesign = JRXmlLoader.load(inputStream);
            JasperReport jasperReport = JasperCompileManager.compileReport(jasperDesign);

            Map<String, Object> parameters = new HashMap();


            JasperPrint jasperPrint = JasperFillManager.fillReport(jasperReport, parameters,
                    new JRBeanCollectionDataSource(exerciseItems));

            final String extension = "jpg";
            final float zoom = 1f;

            if (jasperPrint.getPages().size() == 0) {
                return null;
            }

            try(ByteArrayOutputStream out = new ByteArrayOutputStream()){
                BufferedImage image = (BufferedImage) JasperPrintManager.printPageToImage(jasperPrint, 0,zoom);
                ImageIO.write(image, extension, out);

                return out;
            } catch (IOException e) {
                e.printStackTrace();
            }

        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }
}
