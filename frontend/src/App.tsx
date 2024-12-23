import { useEffect, useState } from 'react';
import { Outlet } from 'react-router';

function App() {
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(false);
  }, []);

  return loading ? null : <Outlet />;
}

export default App;
