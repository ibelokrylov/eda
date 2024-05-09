import * as countries from 'i18n-iso-countries';

const defaultLocale = new Set(['en', 'ru']);

export const COUNTRY_LIST = (locales?: string[]) => {
  try {
    if (locales) {
      for (const locale of locales) {
        if (!defaultLocale.has(locale)) {
          defaultLocale.add(locale);
        }
      }
    }
    defaultLocale.forEach((locale) => {
      countries.registerLocale(
        require(`i18n-iso-countries/langs/${locale}.json`),
      );
    });
  } catch (error) {
    console.error(error);
  }
  return {
    data: (locale: string) =>
      countries.getNames(locale, { select: 'official' }),
    locales: defaultLocale,
    iso: countries.getAlpha2Codes(),
  };
};
