package com.caa.model;

import lombok.Data;
import org.springframework.boot.autoconfigure.domain.EntityScan;

import javax.persistence.*;

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

	public Person() {
	}
}
