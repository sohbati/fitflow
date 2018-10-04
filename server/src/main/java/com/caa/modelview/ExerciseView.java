package com.caa.modelview;

import com.caa.model.Exercise;
import com.caa.model.Person;
import lombok.Data;

/**
 * Created by Reza on 05/08/2018.
 */
@Data
public class ExerciseView  {
    public ExerciseView(Exercise e) {
        this.setId(e.getId());
        this.setName(e.getName());
        this.setLatinName(e.getLatinName());
        this.setCode(e.getCode());
    }
    private long id;
    private String name;
    private String latinName;
    private String code;
}
