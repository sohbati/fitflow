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

    @Query("SELECT e FROM Person e WHERE tenantId=?1 AND (e.mobileNumber like %?2%)")
    public List<Person> queryByMobileNumberForTenant(String tenantId, String mobileNumber);

    @Query("SELECT e FROM Person e WHERE tenantId=?1")
    public List<Person> queryAllForTenant(String tenantId);

    @Query("SELECT e FROM Person e WHERE tenantId=?1 AND (e.firstName LIKE %?2% OR e.lastName LIKE %?2% OR e.mobileNumber LIKE %?2%)")
    public List<Person> queryByNameAndFamiliyAndPhoneForTenant(String tenantId, String str);

}