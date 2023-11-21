"use client"; // Error components must be Client Components

import { ClientError } from "graphql-request";
import { GraphQLError } from "graphql/error/GraphQLError";
// import { GraphQLError } from "graphql/error/GraphQLError";
/**
 * Error Handling
 * https://nextjs.org/docs/app/building-your-application/routing/error-handling
 *
 * The error.js (.tsx) file convention allows you to gracefully handle unexpected runtime errors in nested routes.
 *
 * error.js (.tsx) automatically creates a React Error Boundary that wraps a nested child segment or page.js component.
 * The React component exported from the error.js (.tsx) file is used as the fallback component.
 */

/**
 * Catching rendering errors with an error boundary
 * https://react.dev/reference/react/Component#catching-rendering-errors-with-an-error-boundary
 *
 * By default, if your application throws an error during rendering, React will remove its UI from the screen.
 * To prevent this, you can wrap a part of your UI into an error boundary. An error boundary is a special component
 * that lets you display some fallback UI instead of the part that crashed—for example, an error message.
 */
import { useEffect } from "react";

// https://nextjs.org/docs/app/building-your-application/routing/error-handling
export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };

  /**
   * An error component can use the reset() function to prompt the user to attempt to recover from the error.
   * When executed, the function will try to re-render the Error boundary's contents. If successful, the fallback
   * error component is replaced with the result of the re-render.
   */
  reset: () => void;
}) {
  useEffect(() => {
    // on the client side, the type of `error` is general `Error`
    // although graphql-request throws `ClientError`, due to SSR.

    // Log the error to an error reporting service
    console.error(error);
  }, [error]);

  return (
    <div>
      <h2>Something went wrong!</h2>
      <div>{error.message}</div>
      <button
        onClick={
          // Attempt to recover by trying to re-render the segment
          () => reset()
        }
      >
        Try again
      </button>
    </div>
  );
}
