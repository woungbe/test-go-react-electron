


document.onload = () => {
    console.log("실행은 해봐야지 ? ");
}

const onSubmit = () => {
    apiTest();
}

const onPostTest = () => {
    apiPostTest();
}

function apiTest(){
    reqeust('http://localhost:5000/hello' ,{},res => {
        console.log("res : ", res);
    } , e => alert("에러 : ", e));
}

function apiPostTest(name){
    reqeust('http://localhost:5000/post/exmple01',{
        method:'POST',
        headers: {
            "Contect-Type": "aplication/json",
        },
        body: JSON.stringify({
            title:"Test",
            body: "I am testing !",
            userId:1,
            name:name
        })
    }, res => {
        console.log("res : ", res);
    } , e => alert("에러 : ", e));
}



function reqeust(url, option,funcs , error ){
    fetch(url, option)
    // fetch의 기본 응답 결과는 response 객체
    .then((response) => {
    // response.ok는 status가 200~299사이인 경우 true
    if (response.ok) {
        // 응답을 text 형태로 반환
        return response.text();
    }
    throw new Error(`${response.status} 요청을 처리하지 못했어요!`);
    })
    .then((text) => {
    // document.getElementById("testText").innerHTML = text;
        funcs(text);    
    })
    .catch((e) => {
        error(e)
    });
}
