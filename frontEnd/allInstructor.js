async function populateInstructors() {
  let all_instructors = await fetch(
    `http://localhost:5050/retrieve-instructors`
  );
  let all_instructors_response = await all_instructors.json();
  document.getElementById("instructor_table").style.display = "table";
  document.getElementById("instructor_table_head").innerHTML = `
  <tr>
    <th>${"InstructorCode"}</th>
    <th>${"InstructorName"}</th>
    <th>${"Department"}</th>
    <th>${"CourseName"}</th>
  </tr>`;
  document.getElementById("instructor_table_body").innerHTML =
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

async function insertInstructor() {
  window.location.replace("createInstructor.html");
}

function hideInstructors() {
  document.getElementById("instructor_table").style.display = "none";
}
function setbackpage() {
  window.location.replace("index.html");
}
