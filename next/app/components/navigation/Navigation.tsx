import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { AutoPlayButton } from "./AutoPlayButton";
import { NextButton } from "./NextButton";
import { PrevButton } from "./PrevButton";
import { StepDisplay } from "./StepDisplay";

const fragmentDefinition = graphql(/* GraphQL */ `
  fragment Navigation_Fragment on Page {
    step
    nextStep
    prevStep
  }
`);

export interface NavigationProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  autoNextSeconds?: number | null;
}

export const Navigation = (props: NavigationProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <nav>
      {fragment.step && <StepDisplay step={fragment.step} />}
      {fragment.prevStep && (
        <PrevButton href={`?step=${fragment.prevStep}&skipAnimation=true`} />
      )}
      {fragment.nextStep && (
        <AutoPlayButton
          nextStep={fragment.nextStep}
          autoNextSeconds={props.autoNextSeconds}
        />
      )}
      {fragment.nextStep && <NextButton href={`?step=${fragment.nextStep}`} />}
    </nav>
  );
};
