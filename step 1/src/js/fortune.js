$(document).ready(function() {
        console.log("loading.....");


        function loadFortune() {

            $.getJSON("/api/fortunes/", function (fortunes) {

                    console.log(fortunes);

                    let message = "No luck for you today...";

                    if (fortunes.length > 0) {
                        message = fortunes[0].content

                    }

                    $("#message").text(message);
                }
            );

        }

        loadFortune();
});