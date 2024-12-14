import React, { useState, useEffect } from "react";

function ViewNews() {
  const [news, setNews] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("/api/news")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Не удалось получить новость!!!");
        }
        return response.json();
      })
      .then((data) => {
        setNews(data || []);  
        setLoading(false);
      })
      .catch((error) => {
        console.error("Ошибка:", error);
        setError("При получении новостей произошла ошибка.");
        setLoading(false);
      });
  }, []);

  return (
    <div className="view-news-container">
      <h2>Доступные новости</h2>
      {loading ? (
        <p>Загрузка...</p>
      ) : error ? (
        <p>{error}</p>
      ) : (
        <ul className="news-list">
          {news.length > 0 ? (
            news.map((item) => (
              <li key={item.id} className="news-item">
                <h3 className="news-title">ID: {item.id} - {item.title}</h3>
                <h4 className="news-author">Автор: {item.author}</h4>
                <p className="news-content">{item.content}</p>
              </li>
            ))
          ) : (
            <p>Новости недоступны.</p>
          )}
        </ul>
      )}
    </div>
  );
}

export default ViewNews;
