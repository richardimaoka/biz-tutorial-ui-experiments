import { css } from "@emotion/react";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

interface AutoPlayActive {
  typename: "AutoPlayActive";
  timeoutId: number;
}

interface AutoPlayStopped {
  typename: "AutoPlayStopped";
}

type AutoPlayState = AutoPlayActive | AutoPlayStopped;

const stopped: AutoPlayStopped = { typename: "AutoPlayStopped" };

interface AutoPlayButtonProps {
  nextStep: string;
}

export const AutoPlayButton = ({ nextStep }: AutoPlayButtonProps) => {
  const [state, setState] = useState<AutoPlayState>(stopped);
  const router = useRouter();

  const onClick = () => {
    switch (state.typename) {
      case "AutoPlayStopped":
        // schedule autoPlay to next step
        const timeoutId = window.setTimeout(() => {
          router.replace({ query: { ...router.query, step: nextStep } });
        }, 1000);

        setState({ typename: "AutoPlayActive", timeoutId: timeoutId });
        break;
      case "AutoPlayActive":
        // remove schedule autoPlay to next step
        window.clearTimeout(state.timeoutId);
        setState(stopped);
        break;
      default:
        const _exhaustiveCheck: never = state;
        return _exhaustiveCheck;
    }
  };

  const AutoPlayText = (): JSX.Element => {
    switch (state.typename) {
      case "AutoPlayStopped":
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
      case "AutoPlayActive":
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
