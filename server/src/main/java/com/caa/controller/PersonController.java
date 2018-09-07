package com.caa.controller;

import com.caa.dao.PersonDao;
import com.caa.model.Person;
import com.caa.modelview.PersonView;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.*;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@CrossOrigin(origins = "*")
@RestController
@Transactional(isolation= Isolation.READ_COMMITTED)
public class PersonController {

    private static Logger logger = LoggerFactory.getLogger(PersonController.class);
	
    @Autowired
    private PersonDao personDao;

   /* @RequestMapping(value = { "/" }, method = RequestMethod.GET)
    public HttpEntity<String> sanityCheck() {
    	
    	logger.info("SanityCheck entered...");
    	
		// Return some HTML
    	
    	HttpHeaders headers = new HttpHeaders();
    	headers.setContentType(MediaType.TEXT_PLAIN);
    	
    	String msg = "It's alive! It's alive!";
    	
    	return new HttpEntity<String>(msg, headers);	
    }*/

    @RequestMapping(method = RequestMethod.GET, value = "/getPerson/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
    @ResponseBody
    public Person getPerson(@PathVariable("id") long id) {

    	logger.info("getPerson entered: id= " + id);
    	
    	return personDao.findOne(id);
    }
    
	@RequestMapping(value="/getPersons", method = RequestMethod.GET, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public Iterable<Person> getPersons() {
    	
    	logger.info("findAll entered...");
    	
		return personDao.findAll();
	}

	@RequestMapping(value="/getPersonShortList", method = RequestMethod.GET, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public Iterable<PersonView> getPersonShortList() {

    	logger.info("getPersonShortList entered...");

		List<Person> list = personDao.findAll();
		List<PersonView> result = new ArrayList<>();

		for (Person p : list) {
			PersonView pv = new PersonView(p);
			result.add(pv);
		}

		return result;
	}

	@RequestMapping(value="/savePerson", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	@Transactional
	public Person savePerson(@RequestBody Person person) {
    	
    	logger.info("savePerson entered...");
    	personDao.save(person);
		return person;
	}

	@RequestMapping(method = RequestMethod.DELETE, value = "/deletePerson/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
	@Transactional
	@ResponseBody
	public void deletePerson(@PathVariable("id") long id) {
		logger.info("delete entered: id= " + id);
		try {
			personDao.delete(id);
		}catch (Exception e) {
			throw new DataIntegrityViolationException(e.getMessage());
		}

	}

	// Convert a predefined exception to an HTTP Status code
	@ResponseStatus(value= HttpStatus.EXPECTATION_FAILED,
			reason="Data integrity violation")  //
	@ExceptionHandler(DataIntegrityViolationException.class)
	public void conflict() {
		// Nothing to do
	}
}