function userlogin() {
  // const emailId  = document.getElementById("username").value;
  //const password = document.getElementById("password").value;
  const returnvalue = document.getElementById("response");
  //id1.innerHTML = emailId.value+password.value
  //id1.addEventListener('click',userlogin);

  var valis = CheckValidity();
  returnvalue.innerHTML = valis;
  returnvalue.addEventListener("click", CheckValidity);
  return valis[0];
}

async function CheckValidity() {
  //const endpoint = new URL(`http://localhost:5050/instructor-login/${username}/${password}`);
  
  // const endpoint = new URL ('http://localhost:5050/retrieve-instructors')
  let response = await fetch("http://localhost:5050/retrieve-instructors");
  let json;
  if (response.ok === true) {
    json = await response.json();
    console.log(json[0]);
  } else {
    console.log("iam here");
    alert("HTTP-Error: " + response.status);
  }
  let raw = JSON.stringify(json,undefined,2);
  return raw;
}
