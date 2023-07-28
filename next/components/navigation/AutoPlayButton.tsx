import { css } from "@emotion/react";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

interface Scheduled {
  kind: "Scheduled";
  timeoutId: number;
}

interface Transitioned {
  kind: "Transitioned";
  step: string;
}

interface Stopped {
  kind: "Stopped";
}

type AutoPlayState = Scheduled | Transitioned | Stopped;

interface AutoPlayButtonProps {
  nextStep: string;
}

export const AutoPlayButton = ({ nextStep }: AutoPlayButtonProps) => {
  const [state, setState] = useState<AutoPlayState>({ kind: "Stopped" });
  const router = useRouter();

  // effectful code
  useEffect(() => {
    switch (state.kind) {
      case "Scheduled":
        break; // do nothing
      case "Transitioned":
        // This is an important state as React re-renders this component *BEFORE* updating the URL query string
        // without the Transitioned state, setTimeout is called twice for the same nextStep
        if (state.step !== nextStep) {
          const tid = window.setTimeout(() => {
            router.replace({
              query: { ...router.query, step: nextStep },
            });
            setState({ kind: "Transitioned", step: nextStep });
          }, 1000);
          setState({ kind: "Scheduled", timeoutId: tid });
        }
        break;
      case "Stopped":
        break; // do nothing
      default:
        const _exhaustiveCheck: never = state;
        return _exhaustiveCheck;
    }
  }, [state, nextStep, router]);

  const onClick = () => {
    switch (state.kind) {
      case "Scheduled":
        window.clearTimeout(state.timeoutId);
        setState({ kind: "Stopped" });
        break;
      case "Transitioned":
        setState({ kind: "Stopped" });
        break;
      case "Stopped":
        const tid = window.setTimeout(() => {
          router.replace({
            query: { ...router.query, step: nextStep },
          });
          setState({ kind: "Transitioned", step: nextStep });
        }, 1000);
        setState({ kind: "Scheduled", timeoutId: tid });
        break;
      default:
        const _exhaustiveCheck: never = state;
        return _exhaustiveCheck;
    }
  };

  const AutoPlayText = (): JSX.Element => {
    switch (state.kind) {
      case "Scheduled":
        return (
          <div
            css={css`
              font-size: 16px;
              height: 18px;
            `}
          >
            Stop AutoPlay
          </div>
        );
      case "Transitioned":
        return (
          <div
            css={css`
              font-size: 16px;
              height: 18px;
            `}
          >
            Stop AutoPlay
          </div>
        );
      case "Stopped":
        return (
          <>
            <div
              css={css`
                font-size: 16px;
                height: 18px;
              `}
            >
              AutoPlay
            </div>
            <div
              css={css`
                font-size: 8px;
                line-height: 8px;
              `}
            >
              to next milestone
            </div>
          </>
        );
    }
  };

  return (
    <button
      css={css`
        //positioning
        position: fixed;
        bottom: 0px;
        left: 50%;
        transform: translate(-50%, 0%);

        //sizing
        width: 120px;
        height: 40px;

        // color and styles
        background-color: rgba(255, 255, 255, 0.8);
        color: black;
        border-style: none;
      `}
      onClick={onClick}
    >
      <AutoPlayText />
    </button>
  );
};
