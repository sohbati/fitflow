package com.caa.controller;

import com.caa.model.Person;
import com.caa.model.Program;
import com.caa.model.ProgramExerciseItem;
import com.caa.modelview.ImageView;
import com.caa.modelview.ProgramExerciseItemView;
import com.caa.modelview.ProgramView;
import com.caa.report.ExportReport;
import com.caa.report.ProgramExercisesReportDTO;
import com.caa.services.ExerciseService;
import com.caa.services.ProgramExerciseItemService;
import com.caa.services.ProgramService;
import com.caa.util.DateUtil;
import com.caa.util.ImageUtil;
import org.apache.commons.io.IOUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.util.StreamUtils;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletResponse;
import java.awt.*;
import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.List;
@CrossOrigin(origins = "*")
@RestController
public class ProgramController {

    private static Logger logger = LoggerFactory.getLogger(ProgramController.class);

    @Autowired
	ExerciseService exerciseService;

    @Autowired
	ProgramExerciseItemService programExerciseItemService;

    @Autowired
	ProgramService programService;


    @RequestMapping(method = RequestMethod.GET, value = "/getProgram/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
    @ResponseBody
    public Program getProgram(@PathVariable("id") long id) {
    	return programService.findOne(id);
    }

	@RequestMapping(value="/getPrograms/{personId}")
	@ResponseBody
	public Iterable<ProgramView> getPersonsPrograms(@PathVariable("personId") long personId) {
		return programService.getPersonsPrograms(personId);
	}

	@RequestMapping(value="/saveProgram", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public ProgramView saveProgram(@RequestBody ProgramView programView) {
    	return programService.saveProgram(programView);
	}

	@RequestMapping(method = RequestMethod.DELETE, value = "/deleteProgram/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public void deleteProgram(@PathVariable("id") long id) {
		logger.info("delete entered: id= " + id);
		programService.deleteProgram(id);
	}

	@RequestMapping(method = RequestMethod.GET, value = "/getProgramExerciseImage/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public ImageView getProgramexerciseImage(@PathVariable("id") long id) {
		return programService.getProgramexerciseImage(id);
	}

	@RequestMapping(value = "/shareProgramImage/{id:[\\d]+}", method = RequestMethod.GET,
			produces = MediaType.IMAGE_JPEG_VALUE)
	public void getProgramExerciseListAsImage(HttpServletResponse response, @PathVariable("id") long id) throws IOException {
		 byte[] bytes = programService.getProgramExerciseListAsImage(id);

		response.setContentType(MediaType.IMAGE_JPEG_VALUE);
		StreamUtils.copy(bytes, response.getOutputStream());
	}

	@RequestMapping(value = "/getProgramExercisePDF/{id:[\\d]+}", method = RequestMethod.GET,
			produces = MediaType.APPLICATION_PDF_VALUE)
	public ResponseEntity<byte[]>  getProgramExerciseListAsPDF(@PathVariable("id") long id) throws IOException {
		 byte[] bytes = programService.getProgramExerciseListAsPDF(id);

		HttpHeaders headers = new HttpHeaders();
		headers.setContentType(MediaType.parseMediaType("application/pdf"));
		// Here you have to set the actual filename of your pdf
		String filename = "output.pdf";
		headers.setContentDispositionFormData(filename, filename);
		headers.setCacheControl("must-revalidate, post-check=0, pre-check=0");
		ResponseEntity<byte[]> response = new ResponseEntity<>(bytes, headers, HttpStatus.OK);
		return response;
    }

//	@RequestMapping(value = "/getProgramExercisePDF/{id:[\\d]+}", method = RequestMethod.GET,
//			produces = MediaType.APPLICATION_PDF_VALUE)
//	public void getProgramExerciseListAsPDF(HttpServletResponse response, @PathVariable("id") long id) throws IOException {
//		 byte[] bytes = programService.getProgramExerciseListAsPDF(id);
//
//		response.setContentType(MediaType.APPLICATION_PDF_VALUE);
//		StreamUtils.copy(bytes, response.getOutputStream());
//	}

	@RequestMapping(method = RequestMethod.GET, value = "/getPersonProgramsAllSizes/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public String[][] getPersonProgramsAllSizes(@PathVariable("id") long id) {
		logger.info("getPersonAllSizes entered: id= " + id);
		return programService.getPersonAllSizes(id);
	}


	@PostMapping(value = "/uploadProgramPictures/{personMobileNumber}/{programId}/{imageName}", consumes = "multipart/form-data")
	public void uploadProgramPictures(@PathVariable("personMobileNumber") String personMobileNumber,
									  @PathVariable("programId") long id,
									@PathVariable("imageName") String imageName,
									@RequestPart("pic") MultipartFile picture) throws IOException {
		byte[] img = ImageUtil.handleFileUpload(picture);
		String type = ImageUtil.getImageFileSuffixTypeByContentType(picture.getContentType());
		programService.saveProgramPicture(img, personMobileNumber, id, imageName, type);
	}

	@GetMapping(value = "/downloadProgramPicturesOriginal/{personMobileNumber}/{programId}/{imageNumber}")
	public void downloadProgramPicturesOriginal(HttpServletResponse response, @PathVariable("personMobileNumber") String personMobileNumber,
									  @PathVariable("programId") long programId,
									@PathVariable("imageNumber") int imageNumber) throws IOException {

		byte[] img = programService.loadProgramPictureOriginal( personMobileNumber, programId, imageNumber);

		if (img.length == 0) {
			img = ImageUtil.loadImageInbytes("dummy", "dummy-body.png");
		}
		InputStream in = new ByteArrayInputStream(img);

		response.setContentType(MediaType.IMAGE_JPEG_VALUE);
		//org.apache.commons.io
		IOUtils.copy(in, response.getOutputStream());

	}

	@GetMapping(value = "/downloadProgramPicturesShrinked/{personMobileNumber}/{programId}/{imageNumber}")
	public void downloadProgramPicturesSmall(HttpServletResponse response, @PathVariable("personMobileNumber") String personMobileNumber,
									  @PathVariable("programId") long programId,
									@PathVariable("imageNumber") int imageNumber) throws IOException {

		byte[] img = programService.loadProgramPictureShrinked( personMobileNumber, programId, imageNumber);

		if (img.length == 0) {
			img = ImageUtil.loadImageInbytes("dummy", "dummy-body.png");
		}
		InputStream in = new ByteArrayInputStream(img);

		response.setContentType(MediaType.IMAGE_JPEG_VALUE);
		//org.apache.commons.io
		IOUtils.copy(in, response.getOutputStream());

	}
}