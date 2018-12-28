package com.caa.services;

import com.caa.dao.PersonDao;
import com.caa.model.Person;
import com.caa.modelview.PersonView;
import com.caa.services.security.impl.CustomUserDetailsService;
import com.caa.util.DateUtil;
import com.caa.util.ImageUtil;
import org.apache.commons.lang.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;

import java.io.IOException;
import java.util.List;

/**
 * Created by Reza on 13/09/2018.
 */
@Repository
@Transactional(isolation= Isolation.READ_COMMITTED)
public class PersonService {

    public static String PERSON_MAIN_PICTURE = "personPicture";

    @Autowired
    PersonDao personDao;

    @Autowired
    ImageUtil imageService;

    @Autowired
    TenantConfigurationService tenantConfigurationService;

    public Person findOne(long id) {
      return personDao.findOne(id);
    }

    public List<Person> findAllForTenant() {
        String tenantId = CustomUserDetailsService.getCurrentUserTenant();
        return personDao.queryAllForTenant(tenantId);
    }

    public List<Person> findByNameAndFamiliyAndPhone(String str) {
        String tenantId = CustomUserDetailsService.getCurrentUserTenant();
        return personDao.queryByNameAndFamiliyAndPhoneForTenant(tenantId, str);
    }
    @Transactional
    public Person savePerson(PersonView personView, byte[] picture, String imageSuffix) throws IOException {
        String confFolder = tenantConfigurationService.getProjectConfigFolder();
        int w = tenantConfigurationService.getShrinkedImageWidth();
        int h = tenantConfigurationService.getShrinkedImageHeight();

        byte[] shrinkedImage = ImageUtil.shrinkImage(w, h, picture, imageSuffix);
        validatePerson(personView);
        if (picture.length > 0) {
            ImageUtil.saveToFileSystem(confFolder,
                    personView.getMobileNumber(),
                    PERSON_MAIN_PICTURE,
                    imageSuffix,
                    picture);
        }

        Person person = null;
        if (personView.getId() > 0) {
            person = personDao.getOne(personView.getId());
        }else {
            person = new Person();
        }
        Person saveCandidate = toEntity(personView, person, shrinkedImage, imageSuffix);
        personDao.save(saveCandidate);

        return saveCandidate;
    }



    @Transactional
    public Person savePerson(PersonView personView) {
        validatePerson(personView);
        Person person = null;
        if (personView.getId() > 0) {
            person = personDao.getOne(personView.getId());
        }else {
            person = new Person();
        }

        Person saveCandidate = toEntity(personView, person);
         personDao.save(saveCandidate);
         return saveCandidate;
    }

    private void validatePerson(PersonView personView) {
        long id = personView.getId();
        if (StringUtils.isEmpty(personView.getMobileNumber())) {
            throw new RuntimeException("شماره موبایل اجباری است");
        }

        String tenantId = CustomUserDetailsService.getCurrentUserTenant();
        List<Person> list = personDao.queryByMobileNumberForTenant(tenantId, personView.getMobileNumber());
        if (id == 0) {
            if (list.size() > 0) {
                throw new RuntimeException("شماره موبایل نمی تواند تکراری باشد");
            }
        }else {
            if (list.get(0).getId() != id) {
                throw new RuntimeException("شماره موبایل تغییر داده شده متعلق به فرد دیگری است");
            }
        }
    }

    @Transactional
    public void delete(long id) {
        personDao.delete(id);
    }

    private Person toEntity(PersonView personView, Person person) {
        return  toEntity(personView, person, new byte[0], "");
    }

    private Person toEntity(PersonView personView,Person person, byte[] picture, String imageSuffix) {
        person.setTenantId(CustomUserDetailsService.getCurrentUserTenant());

        person.setId(personView.getId());
        person.setFirstName(personView.getFirstName());
        person.setLastName(personView.getLastName());
        person.setFatherName(personView.getFatherName());
        person.setDisability(personView.getDisability());
        person.setMobileNumber(personView.getMobileNumber());
        person.setAddress(personView.getAddress());
        person.setBirthDate(DateUtil.getGregorianDate(personView.getBirthDate()));
        if (picture.length > 0) {
            person.setShrinkedImage(picture);
            person.setImageSuffix(imageSuffix);
        }
        return person;
    }
}
