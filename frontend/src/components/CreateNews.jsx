import React, { useState } from "react";

const CreateNews = () => {
  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [content, setContent] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    const newsData = { title, author, content };

    try {
      const response = await fetch("/api/news/create", {  
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newsData),
      });

      if (response.ok) {
        alert("Новость успешно создана!");
        setTitle("");
        setAuthor("");
        setContent("");
      } else {
        alert("Не удалось создать новость!!!");
      }
    } catch (error) {
      console.error("Ошибка создания новости:", error);
      alert("При создании новости возникла ошибка!");
    }
  };

  return (
    <div className="form-container">
      <form onSubmit={handleSubmit}>
        <h2>Написать новость</h2>
        <label htmlFor="title">Введите заголовок:</label>
        <input
          type="text"
          id="title"
          name="title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />

        <label htmlFor="author">Введите автора:</label>
        <input
          type="text"
          id="author"
          name="author"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
          required
        />

        <label htmlFor="content">Введите содержимое:</label>
        <textarea
          id="content"
          name="content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          required
        />

        <button type="submit" className="button">
          Создать
        </button>
      </form>
    </div>
  );
};

export default CreateNews;
