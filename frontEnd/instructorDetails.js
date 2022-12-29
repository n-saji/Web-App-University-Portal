async function populateInstructors() {
  let all_instructors = await fetch(
    `http://localhost:5050/retrieve-instructors`
  );
  let all_instructors_response = await all_instructors.json();
  for (let i = 0; i < all_instructors_response.length; i++) {
    let each_value = all_instructors_response[i];
    let table1 = document.getElementById("instructor_table");
    let tr = document.createElement("tr");
    tr.innerHTML = `<td>${each_value.InstructorCode}</td>
       <td>${each_value.InstructorName}</td>
       <td>${each_value.Department}</td>
       <td>${each_value.CourseName}</td>`;
    table1.appendChild(tr);
  }
}
populateInstructors();
function setdashboard() {
  window.location.replace("allinstructor.html");
}
function setbackpage() {
  window.location.replace("createInstructor.html");
}