import React from 'react';
import { Route, Routes, Link } from 'react-router-dom';
import CreateNews from './components/CreateNews';
import UpdateNews from './components/UpdateNews';
import DeleteNews from './components/DeleteNews';
import ViewNews from './components/ViewNews';

function App() {
  return (
    <div>
      <h1>EXPRESS YOURSELF!</h1>
      <h2>лучший новостной портал (ну, почти)</h2>

      <nav>        
        <p><Link to="/view-news" className="button">Показать новости</Link></p>
        <p><Link to="/create-news" className="button">Написать новость</Link></p>
        <p><Link to="/update-news" className="button">Обновить новость</Link></p>
        <p><Link to="/delete-news" className="button">Удалить новость</Link></p>          
        
      </nav>

      <Routes>
        <Route path="/" element={<ViewNews />} />
        <Route path="/view-news" element={<ViewNews />} />
        <Route path="/create-news" element={<CreateNews />} />
        <Route path="/update-news" element={<UpdateNews />} />
        <Route path="/delete-news" element={<DeleteNews />} />        
      </Routes>
    </div>
  );
}

export default App;
