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
    }
    private long id;
    private String name;
    private String latinName;
}
