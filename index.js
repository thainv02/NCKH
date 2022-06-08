var listCoursesBlock = document.querySelector('#list-topic')

var courseAPI = ''


function start(){
    getCourses(function(course) {
        renderCourses (course);
    });
}

start(); 

function getCourses(callback){
    fetch(courseAPI)
        .then(function(response){
            return response.json();
        })
        .then(callback);
}
function renderCourses(){
    
}