package com.caa.model;

import lombok.Data;
import org.springframework.boot.autoconfigure.domain.EntityScan;

import javax.persistence.*;

@Data
@EntityScan
@Entity
@Table(name = "exercise")
public class Exercise {

	@Id
	@Column(name = "id", unique = true, updatable = false, nullable = false)
	@GeneratedValue
	private long id;

	@Column(name = "name", unique = true, updatable = true, insertable = true, nullable = false)
	private String name;

	@Column(name = "latin_name", unique = false, updatable = true, insertable = true, nullable = false)
	private String latinName;

	@Column(name = "involved_muscel", unique = false, updatable = true, insertable = true, nullable = false)
	private String involvedMuscel;

//	@Column(name = "picture", unique = false, updatable = true, insertable = true, nullable = false)
//	private Object picture;

	public Exercise() {
	}
}
