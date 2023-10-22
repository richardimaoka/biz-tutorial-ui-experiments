import { useEffect } from "react";

/**
 * Custm hook to handle window's 'resize' event
 *
 * @return nothing, as it is an effectful hook
 */

export function useResize(
  // Prop to be passed from the parent component
  handleWindowResize: () => void
) {
  useEffect(() => {
    if (window) {
      window.addEventListener("resize", handleWindowResize);
      return () => window.removeEventListener("resize", handleWindowResize);
    }
  }, [handleWindowResize]);
}
