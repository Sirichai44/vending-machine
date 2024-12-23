import { NavLink, useRouteError } from 'react-router-dom';
import AndroidRoundedIcon from '@mui/icons-material/AndroidRounded';
import { IErrorRoute } from '@/store/typings/root';
import { Tooltip } from '@mui/joy';
export default function ErrorPage() {
  const error = useRouteError() as IErrorRoute;
  console.error(error);

  return (
    <div id="error-page" className="flex flex-col items-center justify-center h-screen text-lg">
      <Tooltip title="Go home page" placement="top" arrow>
        <NavLink to="/">
          <AndroidRoundedIcon style={{ fontSize: 50 }} />
        </NavLink>
      </Tooltip>
      <h1>Oops!</h1>
      <p>Sorry, an unexpected error has occurred.</p>
      <p>
        <i>{error?.error?.message || 'An error occurred.'}</i>
      </p>
    </div>
  );
}
