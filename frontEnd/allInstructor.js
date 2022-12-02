async function populateInstructors() {
  let all_instructors = await fetch(
    `http://localhost:5050/retrieve-instructors`
  );
  let all_instructors_response = await all_instructors.json();

  document.getElementById("instructorTable").innerHTML =
    all_instructors_response
      .map(
        (user) =>
          `<tr>
    <td>${user.InstructorCode}</td>
    <td>${user.InstructorName}</td>
    <td>${user.Department}</td>
    <td>${user.CourseName}</td>
    </tr>`
      )
      .join("");
}

async function insertInstructor(){
  let createInstructor = await fetch()
}
