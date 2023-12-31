
interface InputLabelProps {
    label: string;
    type: string;
    id: string;
    onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
    value?: string;
}

function InputLabel(props: InputLabelProps) {
    return (
        <div className="">
            <label
                htmlFor={props.id}
                className="text-gray-500 text-sm "
            >
                {props.label}
            </label>
            <div className="w-full">
                <input
                    id={props.id}
                    className="w-full border border-primary-blue-dark rounded-md px-2 py-1 text-base font-normal text-gray-900"
                    type={props.type}
                    onChange={props.onChange}
                    defaultValue={props.value} />
            </div>
        </div>
    )
}

export default InputLabel