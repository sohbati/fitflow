package com.caa.model;

import lombok.Data;
import org.springframework.boot.autoconfigure.domain.EntityScan;

import javax.persistence.*;
import java.sql.Blob;
import java.util.Date;

@Data
@EntityScan
@Entity
@Table(name = "person")
public class Person {

	@Id
	@Column(name = "id", unique = true, updatable = false, nullable = false)
	@GeneratedValue
	private long id;

	@Column(name = "first_name", unique = true, updatable = true, insertable = true, nullable = false)
	private String firstName;

	@Column(name = "last_name", unique = false, updatable = true, insertable = true, nullable = false)
	private String lastName;

	@Column(name = "mobile_number", unique = false, updatable = true, insertable = true, nullable = false)
	private String mobileNumber;

	// Getters and setters

	@Column(name = "address", unique = false, updatable = true, insertable = true, nullable = true)
	private String address;

	@Column(name = "father_name", unique = false, updatable = true, insertable = true, nullable = true)
	private String fatherName;

	@Column(name = "disability", unique = false, updatable = true, insertable = true, nullable = true)
	private String disability;

	@Column(name = "birth_date", unique = false, updatable = true, insertable = true, nullable = true)
	private Date birthDate;


	@Lob
	@Column(name = "shrinked_image", unique = false, updatable = true, insertable = true, nullable = true)
	private byte[] shrinkedImage;

	@Column(name = "image_suffix", unique = false, updatable = true, insertable = true, nullable = true)
	private String imageSuffix;

	public Person() {
	}
}
