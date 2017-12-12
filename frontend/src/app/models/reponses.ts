export class SuccessMessage {
    constructor(public status: string = "", public message: string = "") {}
}

export class ErrorMessageBody {
    constructor(public name: string = "", public description: string = "") {}
}

export class ErrorMessage {
    constructor(public status: string = "", public error: ErrorMessageBody = new ErrorMessageBody()) {}
}
