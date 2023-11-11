interface Props {
  command: string;
}

export function CommandStringStatic(props: Props) {
  return <code>{props.command}</code>;
}
