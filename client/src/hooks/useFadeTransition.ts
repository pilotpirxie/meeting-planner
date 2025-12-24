import { useEffect, useState } from "react";

export const useFadeTransition = (isVisible: boolean, duration = 150) => {
  const [shouldRender, setShouldRender] = useState(false);
  const [isAnimating, setIsAnimating] = useState(false);

  useEffect(() => {
    if (isVisible) {
      setShouldRender(true);
      setTimeout(() => {
        setIsAnimating(true);
      }, 10);
    } else {
      setIsAnimating(false);
      const timer = setTimeout(() => {
        setShouldRender(false);
      }, duration);
      return () => {
        clearTimeout(timer);
      };
    }
  }, [isVisible, duration]);

  return { shouldRender, isAnimating };
};
