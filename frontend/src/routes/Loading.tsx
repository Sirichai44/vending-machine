import { LinearProgress } from '@mui/joy';

const Loading = () => {
  return (
    <div className="w-full h-screen">
      <LinearProgress thickness={4} variant="plain" />
    </div>
  );
};

export default Loading;
