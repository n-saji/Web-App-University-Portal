function setbackpage() {
  window.location.replace("instructorDetails.html");
}

function setdashboard() {
  window.location.replace("allinstructor.html");
}

async function InsertCourseValues() {
  let course_name = document.getElementById("cn").value;
  let response_for_creation = document.getElementById("response_for_creation");
  if (!course_name) {
    alert("Please enter Course Name.");
    return;
  }
  let createCourse = await fetch(`http://localhost:5050/insert-course`, {
    method: "POST",

    body: JSON.stringify({
      CourseName: course_name,
    }),
  });
  let response = await createCourse.json();
  if (!createCourse.ok) {
    response_for_creation.innerHTML = response + "!!";
  } else {
    console.log(response);
    let response_for_creation = document.getElementById(
      "response_for_creation"
    );
    response_for_creation.innerHTML = "Successffully Created";
  }
}
