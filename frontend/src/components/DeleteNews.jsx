import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const DeleteNews = () => {
  const [newsId, setNewsId] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch(`http://localhost:8080/api/news/delete/${newsId}`, {
        method: "DELETE",
      });

      if (response.ok) {
        alert("Новость успешно удалена!");
        setNewsId("");
        navigate("/");
      } else {
        alert("Не удалось удалить новость!!!");
      }
    } catch (error) {
      alert("Ошибка: " + error.message);
    }
  };

  return (
    <div className="form-container">
      <form onSubmit={handleSubmit}>
        <h2>Удалить новость</h2>
        <label htmlFor="news-id">Введите ID:</label>
        <input
          type="text"
          id="news-id"
          name="news-id"
          value={newsId}
          onChange={(e) => setNewsId(e.target.value)}
          required
        />

        <button type="submit" className="button">
          Удалить
        </button>
      </form>
    </div>
  );
};

export default DeleteNews;
