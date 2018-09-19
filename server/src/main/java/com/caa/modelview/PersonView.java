package com.caa.modelview;

import com.caa.model.Person;
import com.caa.services.PersonService;
import com.caa.util.DateUtil;
import com.caa.util.ImageUtil;
import com.caa.util.PublicUtil;
import lombok.Data;

import java.util.Date;

/**
 * Created by Reza on 05/08/2018.
 */
@Data
public class PersonView {
    public PersonView() {

    }
    public PersonView(Person p) {
        this.setId(p.getId());
        this.setFirstName(p.getFirstName());
        this.setLastName(p.getLastName());
        this.setFatherName(p.getFatherName());
        this.setDisability(p.getDisability());
        this.setMobileNumber(p.getMobileNumber());
        this.setAddress(p.getAddress());
        this.setBirthDate(p.getBirthDate() == null ? "" :
                DateUtil.getShamsiDate(p.getBirthDate() ));
        this.setShrinkedImage(ImageUtil.encodeImage(p.getShrinkedImage()));
        this.setOriginalImage("");
        this.setImageSuffix(p.getImageSuffix());
    }
    private long id;
    private String firstName;
    private String lastName;
    private String fatherName;
    private String disability;
    private String mobileNumber;
    private String address;
    private String birthDate;
    private String shrinkedImage;
    private String imageSuffix;
    private String originalImage;
}
