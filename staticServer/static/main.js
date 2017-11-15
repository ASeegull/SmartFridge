(function() {
    //Creates constants for cashed elements
    const GetContentBtn = document.getElementById('getContents');
    const GetRecipesBtn = document.getElementById('getRecipes');

    //Creates host constant
    const host = 'localhost:8080';

    //Adds eventlisteners to buttons
    GetContentBtn.addEventListener('click', getContent);
    GetRecipesBtn.addEventListener('click', getRecipes);

    //Abstract GET request
    function fetchData(query) {
        return fetch(host + query, {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        }).then(res => jsn = res.json())
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

