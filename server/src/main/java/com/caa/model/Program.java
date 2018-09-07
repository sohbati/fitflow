package com.caa.model;

import lombok.Data;
import org.springframework.boot.autoconfigure.domain.EntityScan;

import javax.persistence.*;
import java.util.Date;
import java.util.List;

/**
 * Created by Reza on 04/08/2018.
 */
@Data
@EntityScan
@Entity
@Table(name = "program")
public class Program {
    @Id
    @Column(name = "id", unique = true, updatable = false, nullable = false)
    @GeneratedValue
    private long id;

    @OneToOne()
    @JoinColumn(name = "person_id")
    private Person person;

    @Column(name = "person_id", unique = false, updatable = false, insertable = false, nullable = false)
    private long personId;

    @Column(name = "program_date", unique = false, updatable = true, insertable = true, nullable = false)
    private Date programDate;

    @Column(name = "program_name", unique = false, updatable = true, insertable = true, nullable = false)
    private String programName;


    @Column(name = "person_age", unique = false, updatable = true, insertable = true, nullable = false)
    private int personAge;

    @Column(name = "person_tall", unique = false, updatable = true, insertable = true, nullable = false)
    private int personTall;

    @Column(name = "person_weight", unique = false, updatable = true, insertable = true, nullable = false)
    private double personWeight;

    @Column(name = "person_chest", unique = false, updatable = true, insertable = true, nullable = false)
    private int personChest;

    @Column(name = "person_waist", unique = false, updatable = true, insertable = true, nullable = false)
    private int personWaist;

    @Column(name = "person_abdomen", unique = false, updatable = true, insertable = true, nullable = false)
    private int personAbdomen;

    @Column(name = "person_arm", unique = false, updatable = true, insertable = true, nullable = false)
    private int personArm;

    @Column(name = "person_forearm", unique = false, updatable = true, insertable = true, nullable = false)
    private int personForeArm;

    @Column(name = "person_thigh", unique = false, updatable = true, insertable = true, nullable = false)
    private int personThigh;

    @Column(name = "person_shin", unique = false, updatable = true, insertable = true, nullable = false)
    private int personShin;

    @Column(name = "person_fat_percentage", unique = false, updatable = true, insertable = true, nullable = false)
    private double personFatPercentage;

    @Column(name = "person_fat_weight", unique = false, updatable = true, insertable = true, nullable = false)
    private double personFatWeight;

    @Column(name = "person_muscle_weight", unique = false, updatable = true, insertable = true, nullable = false)
    private double personMuscleWeight;

    @Column(name = "person_score", unique = false, updatable = true, insertable = true, nullable = false)
    private double personScore;

    @Column(name = "description", unique = false, updatable = true, insertable = true, nullable = true)
    private String description;


}
