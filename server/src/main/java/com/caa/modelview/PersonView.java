package com.caa.modelview;

import com.caa.model.Person;
import lombok.Data;

/**
 * Created by Reza on 05/08/2018.
 */
@Data
public class PersonView {
    public PersonView(Person p) {
        this.setId(p.getId());
        this.setFirstName(p.getFirstName());
        this.setLastName(p.getLastName());
        this.setMobileNumber(p.getMobileNumber());
    }
    private long id;
    private String firstName;
    private String lastName;
    private String mobileNumber;
}
