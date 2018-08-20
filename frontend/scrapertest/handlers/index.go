package handlers

import (
	"net/http"
)

// GetIndex prints stuff
func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<link href="https://fonts.googleapis.com/css?family=Crimson+Text|Libre+Baskerville|Libre+Franklin|Lora|Merriweather|Merriweather+Sans|Playfair+Display|Playfair+Display+SC|Roboto|Slabo+27px|Trirong" rel="stylesheet"> 
			<link rel="stylesheet" type="text/css" href="https://code.jquery.com/ui/1.12.1/themes/base/jquery-ui.css">
			<script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
			<script src="https://code.jquery.com/ui/1.12.1/jquery-ui.min.js"></script>

			<style>
			    #chapter-text {
						background: #E7DEC7;
						color: #5D4232;
						box-sizing: border-box;
						width: 100%;
						left: 0%;
						position: relative;
						border: 1px solid #888;
						/*box-shadow: 7px 7px 3px #ccc;*/
						font-size: 18px;
						font-family: 'Libre Baskerville', serif;
						zoom: 100%;
						padding: 5px;
						white-space: pre-line;
						/*overflow: scroll;*/
					}
			</style>
		</head>
		<body>

			<section id="stories">
					<p>
						<input type="url" id="story-url" autocomplete="on" placeholder="wanderinginn.com"></input>
						<button id="add-story">Add Story</button>
					</p>
					<p>
					<select name="story" id="select-story"></select>
					<select id="select-chapter"></select>
					</p>
			</section>

			<pre id="chapter-text">
			Enter a story url above.
			Like this one: https://wanderinginn.com
			</pre>
		<script>
					currentStory = {};
					loadStory = function(id) {
						$.getJSON("story/"+id, function(data) { 
							currentStory = data;
							console.log(data); 
							$("#select-chapter > option").remove();
							for(var chapterId=0; chapterId<data.NumChapters; chapterId++) {
								$('#select-chapter').append($('<option>', {
									value: chapterId,
									text: chapterId
								}));
							}
							loadChapter(0);
						});
					};
					currentChapter = {};
					loadChapter = function(id) {
						$.getJSON("story/"+currentStory.ID+"/chapter/"+id, function(data) {
							currentChapter = data;
							$("#chapter-text").html(data.HTML);
							console.log(data);
						});
					}
					$(document).ready(function() {
						$("#add-story").click(function() { 
							$.post( "story", JSON.stringify({ url: $("#story-url")[0].value }))
								.done(function( data ) {
									$('#select-story').append($('<option>', {
										value: data.ID,
										text: data.Title
									}));
									loadStory(data.ID);
								})
								.fail(function( data ) {
									alert( "Fail: " + data );
								});
						});
						$("#select-story").change(function(){
							loadStory($(this).val());
						});
						$("#select-chapter").change(function(){
							loadChapter($(this).val());
						});
					});
		</script>
		</body>
		</html>
	`))
}
