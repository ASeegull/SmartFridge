(function() {
    //Creates constants for cashed elements
    const GetContentBtn = document.getElementById('getContents');
    const GetRecipesBtn = document.getElementById('getRecipes');

    //Creates host constant
    const host = 'http://localhost:8080';

    //Adds eventlisteners to buttons
    GetContentBtn.addEventListener('click', getContent);
    GetRecipesBtn.addEventListener('click', getRecipes);

    //Abstract GET request
    function fetchData(query) {
        return fetch(query, {
            method: "GET",
            headers: {
                'Accept': 'application/json, text/plain, */*',
                "Content-Type": "application/json, text/plain"
            }
        }).then(res => {
            let contentType = res.headers.get("content-type");
            console.log(contentType);
            if(contentType && contentType.indexOf("application/json") !== -1) {
              return res.json().then(data => {
                console.log(data);
              });
            } else {
              return res.text().then(text => {
                console.log(text);
              });
            }})
          .catch(e => console.log(e));
        }

    //Fetches content of user's fridge from server
    function getContent() {
        let jsn = fetchData('/content');
    }

    //Get's recipes based on content of fridge
    function getRecipes() {
        let jsn = fetchData('/cookbook');
    }
})();

