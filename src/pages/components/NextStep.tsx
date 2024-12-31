interface iNextStep {
  label: string;
  link: string;
  classes: string;
}

/**
 * @param {object} props
 * @param {label} props.string
 * @param {link} props.link
 * @param {classes} props.classes
 */
export default function NextStep(props: iNextStep) {
  return (
    <a href={`${props.link}`}>
      <button className={`${props.classes}`}>{props.label}</button>
    </a>
  );
}
