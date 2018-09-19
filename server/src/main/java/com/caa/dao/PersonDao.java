package com.caa.dao;

import com.caa.model.Person;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import org.springframework.stereotype.Repository;

import javax.transaction.Transactional;
import java.util.List;

@Repository
@RepositoryRestResource
@Transactional
public interface PersonDao extends JpaRepository<Person, Long> {

    public List<Person> findByMobileNumber(String mobileNumber);

    @Query("SELECT e FROM Person e WHERE e.firstName LIKE %?1% OR e.lastName LIKE %?1% OR e.mobileNumber LIKE %?1%")
    public List<Person> findByNameAndFamiliyAndPhone(String str);

}