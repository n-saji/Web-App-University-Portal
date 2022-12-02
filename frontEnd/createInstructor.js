async function InsertInstructorValues() {
  let instructorcode = document.getElementById("ic").value;
  if (!instructorcode) {
    alert("Please enter instructorcode.");
    return;
  }
  let instructorname = document.getElementById("in").value;
  if (!instructorname) {
    alert("Please enter instructorname.");
    return;
  }
  let department = document.getElementById("dp").value;
  if (!department) {
    alert("Please enter department.");
    return;
  }
  let coursename = document.getElementById("cn").value;
  if (!coursename) {
    alert("Please enter coursename.");
    return;
  }
  let createInstructor = await fetch(
    `http://localhost:5050/insert-instructor-details`,
    {
      method: "POST",

      body: JSON.stringify({
        InstructorCode: instructorcode,
        InstructorName: instructorname,
        Department: department,
        CourseName: coursename,
      }),
    }
  );
  let response = await createInstructor.json();
  console.log(response);
  if (createInstructor.ok != true) {
    alert(response.Err);
  } else if (createInstructor.ok == true) {
    document.getElementById("responseBody").innerHTML = "Inserted Data";
  }
}
