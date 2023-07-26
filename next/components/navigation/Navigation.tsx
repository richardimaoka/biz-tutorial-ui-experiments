import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { AutoPlayButton } from "./AutoPlayButton";
import { BackToStart } from "./BackToStart";
import { NextButton } from "./NextButton";
import { PrevButton } from "./PrevButton";
import { StepDisplay } from "./StepDisplay";

const fragmentDefinition = graphql(/* GraphQL */ `
  fragment NavigationFragment on Page {
    step
    nextStep
    prevStep
  }
`);

export interface NavigationProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const Navigation = (props: NavigationProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <nav>
      <BackToStart />
      {fragment.step && <StepDisplay step={fragment.step} />}
      {fragment.prevStep && <PrevButton href={`?step=${fragment.prevStep}`} />}
      {fragment.nextStep && <AutoPlayButton nextStep={fragment.nextStep} />}
      {fragment.nextStep && <NextButton href={`?step=${fragment.nextStep}`} />}
    </nav>
  );
};
