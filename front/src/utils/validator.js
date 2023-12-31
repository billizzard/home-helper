class Validator {
    constructor(field, value) {
        this.field = field
        this.value = value
        this.errors = []
    }

    maxLength(max) {
        if (!this.value) {
            return this;
        }

        if (this.value.length > max) {
            this.errors.push(`Length "${this.field}" is more than "${max}"`)
        }

        return this
    }

    required() {
        if (!this.value) {
            this.errors.push(`Field "${this.field}" is required`)
        }

        return this
    }

    getFirstError() {
        if (this.errors.length > 0) {
            return this.errors[0]
        }

        return null
    }
}

export default Validator