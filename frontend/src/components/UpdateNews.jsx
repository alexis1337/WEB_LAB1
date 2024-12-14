import React, { useState } from "react";

function UpdateNews() {
  const [id, setId] = useState("");
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();

    fetch(`http://localhost:8080/api/news/update/${id}`, { 
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title, content }),
    })
      .then((response) => {
        if (response.ok) {
          alert("Новость успешно обновлена!");
          setId("");
          setTitle("");
          setContent("");
        } else {
          alert("Не удалось обновить новость!!!");
        }
      })
      .catch((error) => {
        console.error("Ошибка:", error);
        alert("При обновлении новостей произошла ошибка.");
      });
  };

  return (
    <div className="form-container">      
      <form onSubmit={handleSubmit}>
      <h2>Обновить новость</h2>
        <div className="form-group">
          <label htmlFor="news-id">ID:</label>
          <input
            type="text"
            id="news-id"
            value={id}
            onChange={(e) => setId(e.target.value)}
            required
            className="input-field"
          />
        </div>
        <div className="form-group">
          <label htmlFor="news-title">Изменённый заголовок:</label>
          <input
            type="text"
            id="news-title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            className="input-field"
          />
        </div>
        <div className="form-group">
          <label htmlFor="news-content">Изменённое содержимое:</label>
          <textarea
            id="news-content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
            className="textarea-field"
          />
        </div>
        <button type="submit" className="button">
          Обновить
        </button>
      </form>
    </div>
  );
}

export default UpdateNews;
