package com.caa.model;

import lombok.Data;
import org.springframework.boot.autoconfigure.domain.EntityScan;

import javax.persistence.*;

/**
 * Created by Reza on 16/08/2018.
 */
@Data
@EntityScan
@Entity
@Table(name = "program_exercise_item")
public class ProgramExerciseItem {
    @Id
    @Column(name = "idpk", unique = true, updatable = false, nullable = false)
    @GeneratedValue
    private long idpk;

    @Column(name = "id", unique = false, updatable = true, nullable = false)
    //@GeneratedValue
    private long id;

    @Column(name = "program_id", unique = false, updatable = true, insertable = true, nullable = false)
    private long programId;

    @Column(name = "exercise_id", unique = false, updatable = true, insertable = true, nullable = false)
    private long exerciseId;

    @Column(name = "sub_exercise_id", unique = false, updatable = true, insertable = true, nullable = false)
    private int subExerciseId;

    @Column(name = "exercise_set", unique = false, updatable = true, insertable = true, nullable = false)
    private int exerciseSet;

    @Column(name = "exercise_repeat", unique = false, updatable = true, insertable = true, nullable = false)
    private int exerciseRepeat;

    @Column(name = "exercise_repeat_type", unique = false, updatable = true, insertable = true, nullable = false)
    private String exerciseRepeatType;

    @Column(name = "description", unique = false, updatable = true, insertable = true, nullable = true)
    private String description;
}
