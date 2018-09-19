package com.caa.controller;

import com.caa.dao.PersonDao;
import com.caa.model.Person;
import com.caa.modelview.PersonView;
import com.caa.services.PersonService;
import com.caa.util.ImageUtil;
import com.caa.util.PublicUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.List;

@CrossOrigin(origins = "*")
@RestController
public class PersonController {

    private static Logger logger = LoggerFactory.getLogger(PersonController.class);

    @Autowired 	PersonService personService;

	@RequestMapping(value="/getPersons", method = RequestMethod.GET, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public Iterable<Person> getPersons() {
    	
    	logger.info("findAll entered...");
    	
		return personService.findAll();
	}

	@RequestMapping(value="/getPersonShortList", method = RequestMethod.GET, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public Iterable<PersonView> getPersonShortList() {

    	logger.info("getPersonShortList entered...");

		List<Person> list = personService.findAll();
		List<PersonView> result = new ArrayList<>();

		for (Person p : list) {
			PersonView pv = new PersonView(p);
			result.add(pv);
		}

		return result;
	}

	@PostMapping(value = "/savePersonWithOutImage", consumes = "multipart/form-data")
	@ResponseBody
	public PersonView savePersonWithOutImage(@RequestPart("person") PersonView personView) {
    	
    	logger.info("savePerson entered...");
    	Person person = personService.savePerson(personView);
		return new PersonView(person);
	}

	@PostMapping(value = "/savePersonWithImage", consumes = "multipart/form-data")
	public PersonView savePersonWithImage(@RequestPart("person") PersonView personView,
									  @RequestPart("pic") MultipartFile picture) throws IOException {

    	byte[] img = ImageUtil.handleFileUpload(picture);
		Person person = personService.savePerson(personView, img,
			ImageUtil.getImageFileSuffixTypeByContentType(picture.getContentType()));

		return new PersonView(person);
	}

	@RequestMapping(method = RequestMethod.DELETE, value = "/deletePerson/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public void deletePerson(@PathVariable("id") long id) {
		logger.info("delete entered: id= " + id);
		try {
			personService.delete(id);
		}catch (Exception e) {
			throw new DataIntegrityViolationException(e.getMessage());
		}

	}

	@RequestMapping(method = RequestMethod.GET, value = "/getPerson/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public PersonView getPerson(@PathVariable("id") long id) throws IOException {
		Person person = personService.findOne(id);
		PersonView personView = new PersonView(person);
		personView.setOriginalImage(loadOriginalPersonImage(person));
		return personView;
	}

	//@RequestMapping(method = RequestMethod.GET, value = "/findByNameAndFamilyAndPhone/{searchStr}", produces = MediaType.APPLICATION_JSON_VALUE)
    @RequestMapping(value="/findByNameAndFamilyAndPhone/{searchStr}")
    @ResponseBody
	public List<PersonView> findByNameFamilyPhone(@PathVariable("searchStr") String searchStr) throws IOException {
        List<Person> persons = null;
        if (searchStr == null || searchStr.length() == 0 || "EMPTY".equals(searchStr)) {
	        persons = personService.findAll();
        }else {
            persons = personService.findByNameAndFamiliyAndPhone(searchStr);
        }

		List<PersonView> resultList = new ArrayList<>();
		if (persons == null || persons.size() == 0) {
		    return resultList;
        }
        persons.stream().forEach(person -> {
            PersonView personView = new PersonView(person);
            resultList.add(personView);
        });
		return resultList;
	}
	private String loadOriginalPersonImage(Person p) throws IOException {
		String fileFullPath = PublicUtil.getProjectConfigFolder() + "/" +
				p.getMobileNumber() + "/" + PersonService.PERSON_MAIN_PICTURE + "." + p.getImageSuffix();
		File f = new File(fileFullPath);
		if (!f.exists()) {
			return "";
		}
		byte[] bytes = Files.readAllBytes(f.toPath());
		return ImageUtil.encodeImage(bytes);
	}

	// Convert a predefined exception to an HTTP Status code
	@ResponseStatus(value= HttpStatus.EXPECTATION_FAILED,
			reason="Data integrity violation")  //
	@ExceptionHandler(DataIntegrityViolationException.class)
	public void conflict() {
		// Nothing to do
	}
}