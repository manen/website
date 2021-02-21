function to(url) {
  fetch(url)
    .then(data => data.text())
    .then(data => {
      const newDoc = document.open("text/html", "replace");
      newDoc.write(data);
      newDoc.close();

      history.pushState({ ohHey: "the web is a mess B)", url }, null, url);
    });
}
